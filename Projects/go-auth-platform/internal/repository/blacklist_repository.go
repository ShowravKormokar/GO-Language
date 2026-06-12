package repository

import (
	"context"
	"go-auth-platform/internal/models"

	"gorm.io/gorm"
)

type blacklistRepository struct {
	db *gorm.DB
}

func NewBlacklistRepository(db *gorm.DB) BlacklistRepository {
	return &blacklistRepository{
		db: db,
	}
}

func (r *blacklistRepository) Create(ctx context.Context, token *models.BlacklistedToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}
func (r *blacklistRepository) ExistsByJTI(ctx context.Context, jti string) (bool, error) {
	var count int64

	err := r.db.WithContext(ctx).Model(&models.BlacklistedToken{}).Where("jti=?", jti).Count(&count).Error
	return count > 0, err
}
func (r *blacklistRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).Where("expires_at<NOW()").Delete(&models.BlacklistedToken{}).Error
}
