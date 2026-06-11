package migrations

import (
	"fmt"
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/models"
)

func RunMigrations() error {
	err := config.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.RefreshToken{},
		&models.BlacklistedToken{},
	)

	if err != nil {
		return err
	}

	fmt.Println("Migration Completed")

	return nil
}
