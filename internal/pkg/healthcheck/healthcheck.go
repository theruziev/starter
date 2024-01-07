package healthcheck

import (
	"net/http"

	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewHealthCheck() *health.Server {
	healthcheck := health.NewServer()
	return healthcheck
}

func NewLivenessHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

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
