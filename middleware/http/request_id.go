package httpmiddleware

import (
	ctxmeta "go/helperpkg/ctx_meta"
	"go/helperpkg/middleware"
	"net/http"

	"github.com/google/uuid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(middleware.RequestIDHeader)
		if id == "" {
			id = uuid.NewString()
		}

		ctx := ctxmeta.WithRequestID(r.Context(), id)

		w.Header().Set(middleware.RequestIDHeader, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
