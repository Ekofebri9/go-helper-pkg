package errors

type AppError struct {
	Type    Type
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

// Unwrap for errors.Is / errors.As
func (e *AppError) Unwrap() error {
	return e.Err
}

func New(t Type, code, message string) *AppError {
	return &AppError{
		Type:    t,
		Code:    code,
		Message: message,
	}
}

func Wrap(t Type, code, message string, err error) *AppError {
	return &AppError{
		Type:    t,
		Code:    code,
		Message: message,
		Err:     err,
	}
}
