package repository

import (
	"context"
	"go-auth-platform/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Preload("Role").Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Preload("Role").First(&user, "id=?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, "id=?", id).Error
}

func (r *userRepository) List(ctx context.Context, page int, limit int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.
		WithContext(ctx).
		Model(&models.User{})

	if search != "" {

		query = query.Where(
			"name ILIKE ? OR email ILIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	err := query.Count(&total).Error

	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	err = query.
		Preload("Role").
		Limit(limit).
		Offset(offset).
		Find(&users).
		Error

	return users, total, err
}
