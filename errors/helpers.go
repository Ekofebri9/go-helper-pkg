package errors

// IsType checks if error is AppError with given type
func IsType(err error, t Type) bool {
	if e, ok := err.(*AppError); ok {
		return e.Type == t
	}
	return false
}

// AsAppError safely cast
func AsAppError(err error) (*AppError, bool) {
	e, ok := err.(*AppError)
	return e, ok
}
