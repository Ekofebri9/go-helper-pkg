package response

type Response interface {
	GetStatusCode() int
}

type Meta struct {
	ResponseCode    int    `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func (m *Meta) GetStatusCode() int {
	return m.ResponseCode
}

type SuccessResponse struct {
	Data any `json:"data,omitempty"`
	Meta
}

type ErrorResponse struct {
	Error any `json:"error"`
	Meta
}

type PaginationResponse struct {
	Data    any `json:"data"`
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
	Meta
}
