package http

import (
	"net/http"

	"go/helperpkg/errors"
	"go/helperpkg/response"
)

func StatusFromError(err *errors.AppError) int {
	switch err.Type {
	case errors.Invalid:
		return http.StatusBadRequest
	case errors.Unauthorized:
		return http.StatusUnauthorized
	case errors.Forbidden:
		return http.StatusForbidden
	case errors.NotFound:
		return http.StatusNotFound
	case errors.Conflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func ErrorResponse(err *errors.AppError) response.Response {
	return response.ErrorResponse(
		err.Code,
		err.Message,
	)
}
