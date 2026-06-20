package auth

import (
	"go-auth-platform/internal/dto/user"
	userDTO "go-auth-platform/internal/dto/user"
)

type LoginResponse struct {
	User user.UserResponse `json:"user"`
}

type LoginResult struct {
	User         userDTO.UserResponse
	AccessToken  string
	RefreshToken string
}
