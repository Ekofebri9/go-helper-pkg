package errors

type Type string

const (
	Invalid      Type = "INVALID"
	Unauthorized Type = "UNAUTHORIZED"
	Forbidden    Type = "FORBIDDEN"
	NotFound     Type = "NOT_FOUND"
	Conflict     Type = "CONFLICT"
	Internal     Type = "INTERNAL"
)
