package errors

// ErrorMessages maps error codes to human-readable messages.
var Messages = map[Code]string{
	InvalidInput:     "The input provided is invalid.",
	NotFound:         "The requested resource was not found.",
	Conflict:         "There was a conflict with the current state of the resource.",
	Unauthorized:     "You are not authorized to perform this action.",
	PermissionDenied: "You do not have permission to access this resource.",
	InternalServer:   "An internal server error occurred. Please try again later.",
}

// GetMessage returns the human-readable message for a given error code.
func GetMessage(code Code) string {
	if msg, exists := Messages[code]; exists {
		return msg
	}
	return "An unknown error occurred."
}
