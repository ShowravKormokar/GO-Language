package service

import (
	"context"
	"errors"
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

// Reset password logic
func (s *PasswordService) ResetPassword(ctx context.Context, token string, newPassword string) error {

	// Hash the given token
	hashToken := utils.HashResetToken(token)

	// Find the token to check validaty
	reset, err := s.resetRepo.FindValidToken(ctx, hashToken)
	if err != nil {
		return errors.New("Invalid or expired token")
	}

	// Check expired or not
	if time.Now().After(reset.ExpiresAt) {
		return errors.New("reset link expired")
	}

	// Find user
	user, err := s.userRepo.FindByID(ctx, reset.UserID)
	if err != nil {
		return err
	}

	// Hash new password
	newPassHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("Something went wrong")
	}

	// Set new password
	user.PasswordHash = newPassHash

	// Update password
	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return errors.New("password not update")
	}

	// Logout from all devices
	err = s.refreshRepo.RevokeByUserID(ctx, user.ID)
	if err != nil {
		return errors.New("service not avialable at this time")
	}

	// Delete token
	err = s.resetRepo.MarkUsed(ctx, reset.ID)
	if err != nil {
		return errors.New("service not avialable at this time")
	}

	return nil
}
