package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gbmoraes-dev/blueprint/internal/config"
)

type Server struct {
	svc *http.Server
	cfg config.Config
}

func New(cfg config.Config, handler http.Handler) *Server {
	return &Server{
		cfg: cfg,
		svc: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	slog.Info("server listening", "addr", s.svc.Addr)

	if err := s.svc.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("server shutting down")

	return s.svc.Shutdown(ctx)
}
