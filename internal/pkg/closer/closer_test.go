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
	require.Equal(t, 1, order[0])
	require.Equal(t, 2, order[1])

}

func TestCloser_Close(t *testing.T) {

}

func TestCloser_Done(t *testing.T) {

}

func TestNewCloser(t *testing.T) {
	closer := NewCloser()
	require.NotNil(t, closer)
}
