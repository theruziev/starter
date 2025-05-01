package healthcheck

import (
	"net/http"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// NewHealthCheck creates a new health check server that can be used
// to report the health status of the application.
func NewHealthCheck() *health.Server {
	healthcheck := health.NewServer()
	return healthcheck
}

// NewLivenessHandler returns an HTTP handler function for the liveness probe.
// The liveness probe is used to determine if the application is running.
// It always returns a 200 OK status code.
func NewLivenessHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// NewReadinessHandler returns an HTTP handler function for the readiness probe.
// The readiness probe is used to determine if the application is ready to accept traffic.
// It checks the health status of the application and returns:
// - 200 OK if the application is ready
// - 503 Service Unavailable if the application is not ready
// - 500 Internal Server Error if there was an error checking the health status
func NewReadinessHandler(server *health.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := server.Check(r.Context(), &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if resp.Status != grpc_health_v1.HealthCheckResponse_SERVING {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
