package closer

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCloser_Add(t *testing.T) {
	c := NewCloser()
	var count int

	c.Add(func(ctx context.Context) error {
		count++
		return nil
	})

	c.Add(func(ctx context.Context) error {
		count += 2
		return nil
	})

	err := c.Close(context.Background())
	require.NoError(t, err)
	require.Equal(t, 3, count)

}

func TestConcurrentAdd(t *testing.T) {
	c := NewCloser()
	var wg sync.WaitGroup

	addFunc := func() {
		c.Add(func(ctx context.Context) error {
			return nil
		})
		wg.Done()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go addFunc()
	}

	wg.Wait()
	err := c.Close(context.Background())
	require.NoError(t, err)
}

func TestCloseOrderAndErrorHandling(t *testing.T) {
	c := NewCloser()
	var order []int

	c.Add(func(ctx context.Context) error {
		order = append(order, 1)
		return errors.New("error 1")
	})

	c.Add(func(ctx context.Context) error {
		order = append(order, 2)
		return nil
	})

	err := c.Close(context.Background())
	require.Error(t, err)
	require.ErrorContains(t, err, "error 1")

	// Check order of execution
	require.Equal(t, 2, order[0])
	require.Equal(t, 1, order[1])

}

func TestCloser_Close(t *testing.T) {
	t.Run("idempotent close", func(t *testing.T) {
		c := NewCloser()
		var count int

		c.Add(func(ctx context.Context) error {
			count++
			return nil
		})

		// First close
		err := c.Close(context.Background())
		require.NoError(t, err)
		require.Equal(t, 1, count)

		// Second close should be a no-op
		err = c.Close(context.Background())
		require.NoError(t, err)
		require.Equal(t, 1, count) // Count should not increase
	})

	t.Run("context cancellation", func(t *testing.T) {
		c := NewCloser()

		c.Add(func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return nil
			}
		})

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel the context before closing

		err := c.Close(ctx)
		require.Error(t, err)
		require.ErrorIs(t, err, context.Canceled)
	})
}

func TestCloser_Done(t *testing.T) {
	c := NewCloser()

	// Done channel should be open before closing
	select {
	case <-c.Done():
		t.Fatal("Done channel should not be closed before Close is called")
	default:
		// Expected behavior
	}

	// Close in a goroutine
	go func() {
		err := c.Close(context.Background())
		require.NoError(t, err)
	}()

	// Done channel should be closed after closing
	<-c.Done() // This should not block if Done is properly closed
}

func TestNewCloser(t *testing.T) {
	closer := NewCloser()
	require.NotNil(t, closer)
}
