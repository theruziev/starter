package healthcheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewHealthCheck(t *testing.T) {
	healthcheck := NewHealthCheck()
	require.NotNil(t, healthcheck)
}
