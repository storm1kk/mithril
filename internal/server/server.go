package server

import (
	"context"
	"github.com/storm1kk/mithril/internal/config"
	"github.com/storm1kk/mithril/internal/handlers"
	"github.com/storm1kk/mithril/internal/healthz"
	"log/slog"
	"net/http"
	"os"
)

type Server struct {
	httpServer *http.Server
	logger     *slog.Logger
}

func NewServer(
	config *config.Config,
	logger *slog.Logger,
) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/management/healthz", healthz.Handler)

	return &Server{
		httpServer: &http.Server{
			Addr:    config.HttpAddress,
			Handler: mux,
		},
		logger: logger,
	}
}

func (s *Server) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(
				"HTTP server could not start listening.",
				slog.String("addr", s.httpServer.Addr),
				slog.Any("error", err),
			)
			os.Exit(1)
		}
	}()
	s.logger.Info("Server started.", slog.String("addr", s.httpServer.Addr))
}

func (s *Server) Shutdown(ctx context.Context) error {

	return s.httpServer.Shutdown(ctx)
}
