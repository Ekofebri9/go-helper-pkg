package response

// Success response with data
func OK(data any) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// Success response with message
func OKMessage(message string) Response {
	return Response{
		Success: true,
		Message: message,
	}
}

// Success response with data + message
func OKWithMessage(data any, message string) Response {
	return Response{
		Success: true,
		Data:    data,
		Message: message,
	}
}

// Error response
func ErrorResponse(code, message string) Response {
	return Response{
		Success: false,
		Error: &Error{
			Code:    code,
			Message: message,
		},
	}
}
