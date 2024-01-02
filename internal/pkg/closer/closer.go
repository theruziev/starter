package closer

import (
	"context"
	"errors"
	"sync"
)

type closer func(ctx context.Context) error

// Closer is a helper to close multiple closers.
type Closer struct {
	closers  []closer
	mu       sync.Mutex
	isDone   chan struct{}
	isClosed bool
	once     sync.Once
}

// NewCloser creates a new Closer.
func NewCloser() *Closer {
	return &Closer{
		closers:  make([]closer, 0),
		isDone:   make(chan struct{}),
		isClosed: false,
	}
}

// Add adds a closer to the Closer.
func (c *Closer) Add(cl closer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.isClosed {
		return
	}
	c.closers = append(c.closers, cl)
}

// Done returns a channel that's closed when Close is called.
func (c *Closer) Done() chan struct{} {
	return c.isDone
}

// Close closes all closers.
func (c *Closer) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isClosed {
		return nil
	}
	c.isClosed = true

	defer func() {
		c.once.Do(func() {
			close(c.isDone)
		})
	}()

	var resultErr []error
	for i := range c.closers {
		fn := c.closers[i]
		if err := fn(ctx); err != nil {
			resultErr = append(resultErr, err)
		}
	}

	return errors.Join(resultErr...)
}
