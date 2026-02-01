package grpc

import (
	"go/helperpkg/errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CodeFromError(err *errors.AppError) codes.Code {
	switch err.Type {
	case errors.Invalid:
		return codes.InvalidArgument
	case errors.Unauthorized:
		return codes.Unauthenticated
	case errors.Forbidden:
		return codes.PermissionDenied
	case errors.NotFound:
		return codes.NotFound
	case errors.Conflict:
		return codes.AlreadyExists
	default:
		return codes.Internal
	}
}

func Error(err *errors.AppError) error {
	st := status.New(CodeFromError(err), err.Message)

	// inject error code (optional tapi recommended)
	st, _ = st.WithDetails(&ErrorDetail{
		Code: err.Code,
	})

	return st.Err()
}

type ErrorDetail struct {
	Code string `json:"code"`
}

func (e *ErrorDetail) Reset()         {}
func (e *ErrorDetail) String() string { return e.Code }
func (e *ErrorDetail) ProtoMessage()  {}
