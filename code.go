package errors

type Code string

const (
	InvalidInput     Code = "invalid_input"
	NotFound         Code = "not_found"
	Conflict         Code = "conflict"
	Unauthorized     Code = "unauthorized"
	PermissionDenied Code = "permission_denied"
	InternalServer   Code = "internal_server_error"
)
