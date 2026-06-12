package dto

type ErrorResponse struct {
	Success bool              `json:"success"` // always false
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}
