package http

import (
	"net/http"

	"go/helperpkg/errors"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func Handler(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			appErr, ok := errors.AsAppError(err)
			if !ok {
				JSON(w, http.StatusInternalServerError, nil)
				return
			}

			status := StatusFromError(appErr)
			JSON(w, status, ErrorResponse(appErr))
		}
	}
}
