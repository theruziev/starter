package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLivenessHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/live", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := NewLivenessHandler()

	handler.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
}

func TestNewReadinessHandler(t *testing.T) {
	t.Run("not serving", func(t *testing.T) {
		healthcheck := NewHealthCheck()
		req, err := http.NewRequest("GET", "/ready", nil)
		require.NoError(t, err)

		healthcheck.Shutdown()
		rr := httptest.NewRecorder()
		handler := NewReadinessHandler(healthcheck)

		handler.ServeHTTP(rr, req)
		require.Equal(t, http.StatusServiceUnavailable, rr.Code)
	})

	t.Run("serving", func(t *testing.T) {
		healthcheck := NewHealthCheck()
		req, err := http.NewRequest("GET", "/ready", nil)
		require.NoError(t, err)

		healthcheck.Resume()
		rr := httptest.NewRecorder()
		handler := NewReadinessHandler(healthcheck)

		handler.ServeHTTP(rr, req)
		require.Equal(t, http.StatusOK, rr.Code)
	})

}
