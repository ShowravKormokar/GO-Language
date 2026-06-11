package migrations

import (
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/constants"
	"go-auth-platform/internal/models"
)

func SeedRoles() error {
	roles := []models.Role{
		{
			ID:          191,
			Name:        constants.RoleAdmin,
			Description: "Full access",
		},
		{
			ID:          282,
			Name:        constants.RoleManager,
			Description: "Manager access",
		},
		{
			ID:          373,
			Name:        constants.RoleEditor,
			Description: "Editor access",
		},
		{
			ID:          464,
			Name:        constants.RoleUser,
			Description: "Normal user",
		},
	}

	for _, role := range roles {
		config.DB.FirstOrCreate(&role, models.Role{Name: role.Name})
	}

	return nil
}
