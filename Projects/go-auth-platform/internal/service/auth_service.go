package service

import (
	"context"
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/constants"
	dto "go-auth-platform/internal/dto/auth"
	"go-auth-platform/internal/models"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/utils"
	"time"
)

type AuthService struct {
	userRepo      repository.UserRepository
	roleRepo      repository.RoleRepository
	refreshRepo   repository.RefreshTokenRepository
	blacklistRepo repository.BlacklistRepository
}

func NewAuthService(
	userRepo repository.UserRepository,
	roleRepo repository.RoleRepository,
	refreshRepo repository.RefreshTokenRepository,
	blacklistRepo repository.BlacklistRepository,
) *AuthService {

	return &AuthService{
		userRepo:      userRepo,
		roleRepo:      roleRepo,
		refreshRepo:   refreshRepo,
		blacklistRepo: blacklistRepo,
	}
}

// Register Service
func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (*models.User, error) {
	// Check email already used or not
	existing, _ := s.userRepo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Get role type [user]
	role, err := s.roleRepo.FindByName(ctx, constants.RoleUser)
	if err != nil {
		return nil, err
	}

	// User model to create new user
	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: hash,
		RoleID:       role.ID,
		IsActive:     true,
	}

	// Create/register new user
	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login sevice
func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResult, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	if !user.IsActive {
		return nil, ErrInactiveUser
	}

	err = utils.CheckPassword(user.PasswordHash, req.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	tokenPair, _, err := utils.GenerateTokenPair(user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return nil, err
	}

	hash := utils.SHA256Hash(tokenPair.RefreshToken)

	refreshToken := &models.RefreshToken{
		UserID:    user.ID,
		TokenHash: hash,
		ExpiresAt: time.Now().Add(config.AppConfig.JWTRefreshTTL),
	}

	err = s.refreshRepo.Create(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	user.LastLoginAt = &now

	_ = s.userRepo.Update(ctx, user)

	return &dto.LoginResult{
		User:         user,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}
