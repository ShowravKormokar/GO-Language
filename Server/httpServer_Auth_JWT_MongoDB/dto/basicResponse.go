package dto

type BasicResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type DataResponse struct {
	BasicResponse
	Data interface{} `json:"data,omitempty"`
}
