package repository

import (
	"context"
	"go-auth-platform/internal/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, page int, limit int, search string) ([]models.User, int64, error)
}

type RoleRepository interface {
	FindByName(ctx context.Context, name string) (*models.Role, error)
	FindAll(ctx context.Context) ([]models.Role, error)
}

type RefreshTokenRepository interface {
	Create(ctx context.Context, token *models.RefreshToken) error
	FindByHash(ctx context.Context, hash string) (*models.RefreshToken, error)
	RevokeByUserID(ctx context.Context, userId uuid.UUID) error
}

type BlacklistRepository interface {
	Create(ctx context.Context, token *models.BlacklistedToken) error
	ExistsByJTI(ctx context.Context, jti string) (bool, error)
	DeleteExpired(ctx context.Context) error
}

type PasswordResetRepository interface {
	Create(ctx context.Context, token *models.PasswordResetToken) error
	FindValidToken(ctx context.Context, hash string) (*models.PasswordResetToken, error)
	MarkUsed(ctx context.Context, id uuid.UUID) error
}
