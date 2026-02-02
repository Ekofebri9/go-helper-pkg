package httpmiddleware

import (
	"fmt"
	"go/helperpkg/logger"
	"net/http"
	"runtime/debug"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(
					r.Context(),
					"panic recovered",
					logger.ErrorAttr(anyToError(err)),
				)

				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func anyToError(v any) error {
	if err, ok := v.(error); ok {
		return err
	}
	return fmt.Errorf("%v\n%s", v, debug.Stack())
}
