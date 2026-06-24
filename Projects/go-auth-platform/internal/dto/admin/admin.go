package admin

import "time"

type AdminUserQuery struct {
	Page     int
	Limit    int
	Search   string
	Role     string
	IsActive *bool
	Sort     string
	Order    string
}

type AdminUserResponse struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Role      RoleResponse `json:"role"`
	IsActive  bool         `json:"is_active"`
	CreatedAt time.Time    `json:"created_at"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}
