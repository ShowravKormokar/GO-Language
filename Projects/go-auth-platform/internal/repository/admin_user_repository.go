package repository

import (
	"context"
	"errors"
	admDto "go-auth-platform/internal/dto/admin"
	"go-auth-platform/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type adminUserRepo struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) *adminUserRepo {
	return &adminUserRepo{db: db}
}

func (r *adminUserRepo) ListUsers(ctx context.Context, query admDto.AdminUserQuery) ([]models.User, int64, error) {
	var user []models.User
	var total int64

	db := r.db.WithContext(ctx).Model(&models.User{}).Preload("Role")

	// Soft deleted users ignore autometically

	// Filter by name or email
	if query.Search != "" {
		search := "%" + query.Search + "%"

		db = db.Where("name ILIKE ? OR email ILIKE ?", search, search)
	}

	// Filter by role id
	if query.Role != "" {
		db = db.Where("role_id IN (?)", r.db.Table("roles").Select("id").Where("name=?", query.Role))
	}

	// Filter by active or inactive user
	if query.IsActive != nil {
		db = db.Where("is_active=?", query.IsActive)
	}

	// Count before pagination
	db.Count(&total)

	// Whitelist to prevent SQL injection
	allowedSort := map[string]bool{
		"created_at": true,
		"name":       true,
		"email":      true,
	}

	if !allowedSort[query.Sort] {
		query.Sort = "created_at"
	}

	order := "DESC"

	if query.Order == "asc" {
		order = "ASC"
	}

	db = db.Order(query.Sort + " " + order)

	// Page limit
	offset := (query.Page - 1) * query.Limit

	err := db.Limit(query.Limit).Offset(offset).Find(&user).Error

	return user, total, err
}

// Find user by ID
func (r *adminUserRepo) FindUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Preload("Role").Where("ID = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	return &user, err
}

// Update user
func (r *adminUserRepo) UpdateFields(ctx context.Context, id uuid.UUID, data map[string]interface{}) error {

	return r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(data).
		Error

}
