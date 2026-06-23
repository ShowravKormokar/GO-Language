package service

import (
	"context"
	"errors"

	urdto "go-auth-platform/internal/dto/user"
	"go-auth-platform/internal/mapper"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/utils"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
}

func NewUserService(userRepo repository.UserRepository, refreshRepo repository.RefreshTokenRepository) *UserService {
	return &UserService{
		userRepo:    userRepo,
		refreshRepo: refreshRepo,
	}
}

// Get current logged in user
func (s *AuthService) GetCurrentUser(ctx context.Context, userId string) (*urdto.UserProfileResponse, error) {
	id := uuid.MustParse(userId)

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	response := mapper.ToUserProfileResponse(user)

	return &response, nil
}

// Change password (Logged in user)
func (s *UserService) ChangePassword(ctx context.Context, userID string, req urdto.ChangePasswordRequest) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	// Find user by ID (Basically current user)
	currentUser, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	// Old password check
	err = utils.CheckPassword(currentUser.PasswordHash, req.CurrentPassword)
	if err != nil {
		return errors.New("current password is incorrect")
	}

	// Hash new password
	newHashPass, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("something went wrong")
	}

	// Assign new password
	currentUser.PasswordHash = newHashPass

	// Update databse (user password update)
	err = s.userRepo.Update(ctx, currentUser)
	if err != nil {
		return err
	}

	// Logout all devices
	err = s.refreshRepo.RevokeByUserID(ctx, currentUser.ID)

	return err
}
