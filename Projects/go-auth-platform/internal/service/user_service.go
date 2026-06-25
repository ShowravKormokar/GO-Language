package service

import (
	"context"
	"errors"

	admDto "go-auth-platform/internal/dto/admin"
	clDto "go-auth-platform/internal/dto/claims"
	urdto "go-auth-platform/internal/dto/user"
	"go-auth-platform/internal/mapper"
	"go-auth-platform/internal/models"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/utils"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo      repository.UserRepository
	roleRepo      repository.RoleRepository
	refreshRepo   repository.RefreshTokenRepository
	blacklistRepo repository.BlacklistRepository
}

func NewUserService(userRepo repository.UserRepository, roleRepo repository.RoleRepository, refreshRepo repository.RefreshTokenRepository, blacklistRepo repository.BlacklistRepository) *UserService {
	return &UserService{
		userRepo:      userRepo,
		roleRepo:      roleRepo,
		refreshRepo:   refreshRepo,
		blacklistRepo: blacklistRepo,
	}
}

// Get current logged in user
func (s *UserService) GetCurrentUser(ctx context.Context, userId string) (*urdto.UserProfileResponse, error) {
	id := uuid.MustParse(userId)

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	response := mapper.ToUserProfileResponse(user)

	return &response, nil
}

// Change password (Logged in user)
func (s *UserService) ChangePassword(ctx context.Context, userID string, claims *clDto.JWTClaims, req urdto.ChangePasswordRequest) error {
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

	// Logout all other devices
	err = s.refreshRepo.RevokeByUserID(ctx, currentUser.ID)
	if err != nil {
		return err
	}

	// blacklist current access token - logout current device
	err = s.blacklistRepo.Create(
		ctx,
		&models.BlacklistedToken{
			JTI:       claims.JTI,
			ExpiresAt: claims.ExpiresAt.Time,
		},
	)

	return err
}

// Assign role (update user)
func (s *UserService) AssignRole(ctx context.Context, targetUserID string, req admDto.AssignRoleRequest) (*models.User, error) {
	userId, err := uuid.Parse(targetUserID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, errors.New("user not found")
	}

	role, err := s.roleRepo.FindByID(ctx, req.RoleID)
	if err != nil {
		return nil, errors.New("role not found")
	}

	// assign
	user.RoleID = role.ID
	user.Role = *role

	err = s.userRepo.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) UpdateUserStatus(ctx context.Context, targetUserID string, req admDto.UpdateUserStatusRequest) (*models.User, error) {
	id, err := uuid.Parse(targetUserID)

	if err != nil {
		return nil, errors.New("invalid user id")
	}

	// find user
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// update status
	user.IsActive = req.IsActive

	err = s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	// If account disabled -> logout all devices
	if !req.IsActive {
		err = s.refreshRepo.RevokeByUserID(ctx, user.ID)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
