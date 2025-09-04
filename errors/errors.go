package errors

import (
	"net/http"
)

type Error struct {
	code    Code
	message any
	error   error
}

func (e *Error) GetCode() Code {
	return e.code
}

func (e *Error) GetMessage() string {
	if msg, ok := e.message.(string); ok {
		return msg
	}
	return ""
}

func (e *Error) GetError() error {
	return e.error
}

func (e *Error) Error() string {
	if e.error != nil {
		return e.error.Error()
	}
	return e.GetMessage()
}

func (e *Error) Unwrap() error {
	return e.error
}

func (e *Error) GetHTTPCode() int {
	switch e.code {
	case InvalidInput:
		return http.StatusBadRequest // 400
	case NotFound:
		return http.StatusNotFound // 404
	case Conflict:
		return http.StatusConflict // 409
	case Unauthorized:
		return http.StatusUnauthorized // 401
	case PermissionDenied:
		return http.StatusForbidden // 403
	case InternalServer:
		return http.StatusInternalServerError // 500
	default:
		return http.StatusInternalServerError // 500
	}
}

func NewError(code Code, message any, err error) *Error {
	return &Error{
		code:    code,
		message: message,
		error:   err,
	}
}

func NewInternalError(err error) *Error {
	return &Error{
		code:    InternalServer,
		message: Messages[InternalServer],
		error:   err,
	}
}

func NewInvalidInputError(message any, err error) *Error {
	return &Error{
		code:    InvalidInput,
		message: message,
		error:   err,
	}
}

func NewNotFoundError(message any, err error) *Error {
	return &Error{
		code:    NotFound,
		message: message,
		error:   err,
	}
}

func NewConflictError(message any, err error) *Error {
	return &Error{
		code:    Conflict,
		message: message,
		error:   err,
	}
}

func NewUnauthorizedError(message any, err error) *Error {
	return &Error{
		code:    Unauthorized,
		message: message,
		error:   err,
	}
}

func NewPermissionDeniedError(message any, err error) *Error {
	return &Error{
		code:    PermissionDenied,
		message: message,
		error:   err,
	}
}
