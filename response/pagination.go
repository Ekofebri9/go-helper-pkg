package response

type PaginationMeta struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

func WithPagination(data any, meta PaginationMeta) Response {
	return Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	}
}
