package repository

import (
	"context"
	"errors"
	"go-auth-platform/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type passwordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) PasswordResetRepository {

	return &passwordResetRepository{
		db: db,
	}
}

// Create a new password reset token in the database
func (r *passwordResetRepository) Create(ctx context.Context, token *models.PasswordResetToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

// Find valid token
func (r *passwordResetRepository) FindValidToken(ctx context.Context, hash string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken

	err := r.db.WithContext(ctx).
		Where("token_hash= ? AND used = ?", hash, false).
		First(&resetToken).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(
				"invalid reset token",
			)
		}
		return nil, err
	}

	return &resetToken, nil
}

// Mark as used
func (r *passwordResetRepository) MarkUsed(ctx context.Context, id uuid.UUID) error {
	err := r.db.WithContext(ctx).
		Model(&models.PasswordResetToken{}).
		Where("id = ?", id).
		Update("used", true).
		Error

	return err
}
