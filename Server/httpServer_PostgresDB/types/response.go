package types

type APIResponse struct {
	Mssg   string      `json:"mssg"`
	Status int         `json:"status"`
	Data   interface{} `json:"omitempty"`
}
