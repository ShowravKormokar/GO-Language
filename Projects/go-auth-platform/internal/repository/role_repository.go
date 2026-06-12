package repository

import (
	"context"
	"go-auth-platform/internal/models"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) FindByName(ctx context.Context, name string) (*models.Role, error) {
	var role models.Role

	err := r.db.WithContext(ctx).Where("name=?", name).First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindAll(ctx context.Context) ([]models.Role, error) {
	var roles []models.Role

	err := r.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}
