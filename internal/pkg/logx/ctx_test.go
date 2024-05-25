package logx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestWithLogger(t *testing.T) {
	ctx := context.Background()
	log := zap.NewExample().Sugar()

	newCtx := WithLogger(ctx, log)

	require.Equal(t, log, newCtx.Value(loggerKey))
}

func TestFromContext(t *testing.T) {
	log := zap.NewExample().Sugar()
	ctx := context.WithValue(context.Background(), loggerKey, log)

	assert.Equal(t, log, FromContext(ctx))

	emptyCtx := context.Background()
	require.Equal(t, DefaultLogger(), FromContext(emptyCtx))
}
