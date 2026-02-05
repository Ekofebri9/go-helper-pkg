package httpmiddleware

import (
	ctxmeta "go/helperpkg/ctx_meta"
	"go/helperpkg/logger"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := ctxmeta.RequestID(r.Context())

		l := logger.Base().With(
			slog.String("request_id", reqID),
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)

		span := trace.SpanFromContext(r.Context())
		if span.SpanContext().IsValid() {
			traceID := span.SpanContext().TraceID().String()
			l = l.With(slog.String("trace_id", traceID))
		}

		ctx := logger.WithContext(r.Context(), l)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
