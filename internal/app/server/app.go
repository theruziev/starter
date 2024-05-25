package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/theruziev/starter/internal/pkg/closer"
	"github.com/theruziev/starter/internal/pkg/healthcheck"
	"github.com/theruziev/starter/internal/pkg/info"
	"github.com/theruziev/starter/internal/pkg/logx"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	opt         *Options
	closer      *closer.Closer
	r           chi.Router
	healthcheck *health.Server

	routerOnce sync.Once
}

type Options struct {
	Addr string
}

func NewServer(cl *closer.Closer, opt *Options) *Server {
	return &Server{
		opt:    opt,
		closer: cl,
	}
}

func (s *Server) Init(ctx context.Context) error {
	s.initHealthCheck(ctx)
	s.initRouter(ctx)
	return nil
}

func (s *Server) initHealthCheck(_ context.Context) {
	h := healthcheck.NewHealthCheck()

	s.healthcheck = h
}

func (s *Server) initRouter(_ context.Context) {
	s.routerOnce.Do(func() {
		r := chi.NewRouter()
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Hello world!"))
		})
		r.Get("/live", healthcheck.NewLivenessHandler())
		r.Get("/ready", healthcheck.NewReadinessHandler(s.healthcheck))
		r.Get("/version", info.Handler())
		s.r = r
	})
}

func (s *Server) Run(ctx context.Context) error {
	logger := logx.FromContext(ctx)
	logger.Infof("server is ready to handle requests at  %s", s.opt.Addr)
	srv := http.Server{
		Addr:              s.opt.Addr,
		Handler:           s.r,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}
	s.closer.Add(func(ctx context.Context) error {
		logger.Warn("server is shutting down")
		if err := srv.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown server: %w", err)
		}

		return nil
	})

	s.healthcheck.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	s.closer.Add(func(ctx context.Context) error {
		logger.Infof("healthcheck is shutting down")
		s.healthcheck.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
		return nil
	})

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}
	return nil
}
