package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/theruziev/starter/internal/app/server"
	"github.com/theruziev/starter/internal/pkg/closer"
	"github.com/theruziev/starter/internal/pkg/logx"
)

type serverCli struct {
	Addr string `help:"Server address" default:":8080" env:"SERVER_ADDR"`
}

func (s *serverCli) Run(cliCtx *contextCli) error {
	logger := logx.NewLogger(cliCtx.LogLevel, cliCtx.IsDebug)
	ctx := logx.WithLogger(context.Background(), logger)
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cl := closer.NewCloser()
	go func() {
		<-ctx.Done()
		logger.Warn("graceful shutdown initiated")
		closeCtx, closeCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer closeCancel()
		if err := cl.Close(closeCtx); err != nil {
			logger.Errorf("failed to close server: %s", err)
		}
	}()

	srv := server.NewServer(cl, &server.Options{
		Addr: s.Addr,
	})

	if err := srv.Init(ctx); err != nil {
		return fmt.Errorf("failed to init server: %w", err)
	}

	if err := srv.Run(ctx); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}

	<-cl.Done()
	return nil
}
