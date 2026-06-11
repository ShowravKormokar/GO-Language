package dto

type APIResponse[T any] struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Data      T      `json:"data,omitempty"`
	RequestID string `json:"request_id,omitempty"`
}
