package service

import (
	"context"
	"go-auth-platform/internal/models"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/utils"
	"time"
)

// Dependency Inject
type PasswordService struct {
	userRepo    repository.UserRepository
	resetRepo   repository.PasswordResetRepository
	refreshRepo repository.RefreshTokenRepository
}

// Constructor
func NewPasswordService(userRepo repository.UserRepository, resetRepo repository.PasswordResetRepository, refreshRepo repository.RefreshTokenRepository) *PasswordService {

	return &PasswordService{
		userRepo:    userRepo,
		resetRepo:   resetRepo,
		refreshRepo: refreshRepo,
	}
}

// Forget password logic
func (s *PasswordService) ForgotPassword(ctx context.Context, email string) (string, error) {

	user, err := s.userRepo.FindByEmail(ctx, email)
	// important
	// don't expose user existence
	if err != nil {
		return "", nil
	}

	token := utils.GenerateResetToken()
	hashToken := utils.HashResetToken(token)

	reset := &models.PasswordResetToken{
		UserID:    user.ID,
		TokenHash: hashToken,
		ExpiresAt: time.Now().Add(60 * time.Second),
	}

	err = s.resetRepo.Create(ctx, reset)
	if err != nil {
		return "", nil
	}

	// dummy email link
	link := "http://localhost:8080/reset-password?token=" + token

	return link, nil
}
