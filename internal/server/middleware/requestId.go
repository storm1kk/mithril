package middleware

import (
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"os"
)

const HeaderKey = "mithril-request-id"

func RequestId(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.NewV7()
		if err != nil {
			logger.Error("Impossible to generate request ID.")
			os.Exit(1)
		}

		// Кладем ID запроса в хедеры, чтобы можно было их достать потом.
		r.Header.Add(HeaderKey, id.String())

		logger.Debug("Request ID generated.", slog.String("id", id.String()))
		next.ServeHTTP(w, r)
	})
}
