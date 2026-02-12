package httpmiddleware

import (
	jwtpkg "go/helperpkg/auth/jwt"
	ctxmeta "go/helperpkg/ctx_meta"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func HTTP[T jwt.Claims](verifier *jwtpkg.Verifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := jwtpkg.ExtractFromHTTP(r)

			claims, err := jwtpkg.Verify[T](verifier, token)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := ctxmeta.WithClaims(r.Context(), claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
