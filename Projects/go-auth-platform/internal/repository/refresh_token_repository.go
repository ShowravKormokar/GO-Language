package repository

import (
	"context"
	"go-auth-platform/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(
	db *gorm.DB,
) RefreshTokenRepository {

	return &refreshTokenRepository{
		db: db,
	}
}

func (r *refreshTokenRepository) Create(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}
func (r *refreshTokenRepository) FindByHash(ctx context.Context, hash string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.WithContext(ctx).Where("token_hash = ? AND revoked = false", hash).First(&token).Error
	if err != nil {
		return nil, err
	}

	return &token, nil
}
func (r *refreshTokenRepository) RevokeByUserID(ctx context.Context, userId uuid.UUID) error {
	return r.db.WithContext(ctx).Model(&models.RefreshToken{}).Where("user_id=?", userId).Update("revoked", true).Error
}
