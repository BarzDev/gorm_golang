package shared_model

type Status struct {
	Code    int
	Message string
}

type PagedResponse struct {
	Status  Status
	Code    int
	Message string
	Data    interface{}
	Paging  Paging
}

type SingleResponse struct {
	Code    int
	Data    interface{}
	Message string
}