package httpmiddleware

import (
	"net/http"
	"time"
)

func Timeout(d time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, d, "request timeout")
	}
}
