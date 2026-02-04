package httpmiddleware

import (
	ctxmeta "go/helperpkg/ctx_meta"
	"go/helperpkg/logger"
	"log/slog"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := ctxmeta.RequestID(r.Context())

		l := logger.Base().With(
			slog.String("request_id", reqID),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		ctx := logger.WithContext(r.Context(), l)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
