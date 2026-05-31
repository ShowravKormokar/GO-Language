package dto

type BasicResponse struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

type DataResponse struct {
	BasicResponse
	Data interface{} `json:"data,omitempty"`
}
