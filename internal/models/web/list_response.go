package web

type ListResponse[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}
