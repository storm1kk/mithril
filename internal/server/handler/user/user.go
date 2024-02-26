package user

import (
	"encoding/json"
	"github.com/storm1kk/mithril/internal/entity"
	"github.com/storm1kk/mithril/internal/server/middleware"
	"github.com/storm1kk/mithril/internal/storage"
	"io"
	"log/slog"
	"net/http"
)

func CreateUser(logger *slog.Logger, storage storage.Storage) http.Handler {
	const op = "handlers.user.createUser"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			// TODO: log
			return // TODO: error method not allowed
		}

		var u entity.User
		// TODO: validation username
		// TODO: validation password
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			logger.Error(
				"Error during decode body.",
				slog.String("op", op),
				slog.Any("error", err),
				slog.String("body", bodyToString(r.Body)),
				slog.String(middleware.HeaderKey, r.Header.Get(middleware.HeaderKey)),
			)
			http.Error(w, err.Error(), http.StatusBadRequest) // TODO: make response JSON
			return
		}

		id, err := storage.CreateUser(u)
		if err != nil {
			logger.Error(
				"Error during create user.",
				slog.String("op", op),
				slog.Any("error", err),
				slog.String(middleware.HeaderKey, r.Header.Get(middleware.HeaderKey)),
			)
			http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: make response JSON
			return
		}

		w.Header().Set("Content-Type", "application/json")
		resp := struct {
			ID       int64
			Username string
		}{
			ID:       id,
			Username: u.Username,
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			logger.Error("Error encoding response.", slog.String("op", op), slog.Any("error", err))
			http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: make response JSON
			return
		}
	})
}

func bodyToString(closer io.ReadCloser) string {
	b, _ := io.ReadAll(closer)
	return string(b)
}
