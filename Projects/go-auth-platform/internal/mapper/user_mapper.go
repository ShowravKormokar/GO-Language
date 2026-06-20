package mapper

import (
	userDTO "go-auth-platform/internal/dto/user"

	"go-auth-platform/internal/models"
)

// Login mapper
func ToUserResponse(u *models.User) userDTO.UserResponse {

	return userDTO.UserResponse{
		ID:          u.ID.String(),
		Name:        u.Name,
		Email:       u.Email,
		Role:        u.Role.Name,
		IsActive:    u.IsActive,
		LastLoginAt: u.LastLoginAt,
		CreatedAt:   u.CreatedAt,
	}
}

// Profile mapper
func ToUserProfileResponse(u *models.User) userDTO.UserProfileResponse {

	return userDTO.UserProfileResponse{
		ID:    u.ID.String(),
		Name:  u.Name,
		Email: u.Email,
		Role: userDTO.RoleResponse{
			ID:          u.Role.ID,
			Name:        u.Role.Name,
			Description: u.Role.Description,
		},

		IsActive:    u.IsActive,
		LastLoginAt: u.LastLoginAt,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
