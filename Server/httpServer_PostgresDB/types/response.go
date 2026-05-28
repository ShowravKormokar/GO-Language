package types

type APIResponse struct {
	Mssg   string      `json:"mssg"`
	Data   interface{} `json:"omitempty"`
	Status int         `json:"status"`
}
