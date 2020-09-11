package http

//Response represent the response of the request
type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//PaginationResponse represent the response for paginated request
type PaginationResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Size    int         `json:"size"`
	Page    int         `json:"page"`
}
