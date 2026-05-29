package types

type APIResponse struct {
	Mssg   string      `json:"mssg"`
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type SuccessResponse struct {
	Mssg   string `json:"mssg"`
	Status int    `json:"status"`
}

type ErrorResponse struct {
	Mssg   string `json:"mssg"`
	Status int    `json:"status"`
	Error  error  `json:"error,omitempty"`
}
