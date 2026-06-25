package admin

import (
	dto "go-auth-platform/internal/dto/user"
	"time"
)

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
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Role      dto.RoleResponse `json:"role"`
	IsActive  bool             `json:"is_active"`
	CreatedAt time.Time        `json:"created_at"`
}

type AssignRoleRequest struct {
	RoleID uint `json:"role_id" validate:"required"`
}
