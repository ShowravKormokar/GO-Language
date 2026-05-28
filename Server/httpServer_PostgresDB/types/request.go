package types

type UserRequest struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
