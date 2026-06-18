package auth

import (
	"go-auth-platform/internal/models"
	"time"
)

type UserResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Role        string     `json:"role"`
	IsActive    bool       `json:"is_active"`
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

type LoginResult struct {
	User         *models.User
	AccessToken  string
	RefreshToken string
}
