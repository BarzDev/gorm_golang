package shared_model

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PagedResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
	Paging Paging      `json:"paging"`
}

type SingleResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
