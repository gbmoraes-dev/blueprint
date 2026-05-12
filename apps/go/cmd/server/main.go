package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gbmoraes-dev/blueprint/internal/config"
	"github.com/gbmoraes-dev/blueprint/internal/logger"
	"github.com/gbmoraes-dev/blueprint/internal/router"
	"github.com/gbmoraes-dev/blueprint/internal/server"
)

func main() {
	cfg := config.Load()

	log := logger.New()
	slog.SetDefault(log)

	srv := server.New(cfg, router.New())

	svcErr := make(chan error, 1)
	go func() {
		svcErr <- srv.Start()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-svcErr:
		log.Error("server failed to start", "error", err)
		os.Exit(1)
	case sig := <-quit:
		log.Info("shutdown signal received", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		if err := srv.Shutdown(ctx); err != nil {
			cancel()
			log.Error("graceful shutdown failed", "error", err)
			os.Exit(1)
		}

		cancel()
		log.Info("server stopped cleanly")
	}
}
