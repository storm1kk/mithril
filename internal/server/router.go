package server

import (
	"github.com/storm1kk/mithril/internal/server/handler/healthz"
	"github.com/storm1kk/mithril/internal/server/handler/root"
	"github.com/storm1kk/mithril/internal/server/handler/user"
	"github.com/storm1kk/mithril/internal/server/middleware"
	"github.com/storm1kk/mithril/internal/storage"
	"log/slog"
	"net/http"
)

func addRoutes(
	mux *http.ServeMux,
	logger *slog.Logger,
	storage storage.Storage,
) {
	mux.Handle("/", defaultMiddlewares(root.Handler(), logger))
	mux.Handle("/management/healthz", defaultMiddlewares(healthz.Handler(), logger))
	mux.Handle("/api/users", defaultMiddlewares(user.CreateUser(logger, storage), logger))
}

func defaultMiddlewares(h http.Handler, logger *slog.Logger) http.Handler {
	return middleware.RequestId(h, logger)
}
