package service

import (
	"context"
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/constants"
	dto "go-auth-platform/internal/dto/auth"
	dtoJWT "go-auth-platform/internal/dto/claims"
	"go-auth-platform/internal/models"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/utils"
	"time"

	"github.com/google/uuid"
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

	// Reload user with Role
	userWithRole, err := s.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return userWithRole, nil
}

// Login sevice
func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResult, error) {

	// Find user by email
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Check user is active
	if !user.IsActive {
		return nil, ErrInactiveUser
	}

	// Compare password correct or not
	err = utils.CheckPassword(user.PasswordHash, req.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Generate token pair
	tokenPair, _, err := utils.GenerateTokenPair(user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return nil, err
	}

	// Hash the token
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

	// Get current time
	now := time.Now()

	// Last login time
	user.LastLoginAt = &now

	_ = s.userRepo.Update(ctx, user)

	return &dto.LoginResult{
		User:         user,
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

// Logout Service
func (s *AuthService) Logout(ctx context.Context, claims *dtoJWT.JWTClaims) error {
	// Create a blacklisted token record
	blacklisted := &models.BlacklistedToken{
		JTI:       claims.JTI,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	// Create a blacklisted token record
	err := s.blacklistRepo.Create(ctx, blacklisted)

	if err != nil {
		return nil
	}

	// Revoke all refresh tokens for the user
	return s.refreshRepo.RevokeByUserID(ctx, uuid.MustParse(claims.UserID))
}

// Refresh token service
func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (*dtoJWT.RefreshResult, error) {
	// Parse the refresh token and validate it
	claims, err := utils.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	// Check if the token is blacklisted
	hash := utils.SHA256Hash(refreshToken)

	// Check if the token is revoked
	stored, err := s.refreshRepo.FindByHash(ctx, hash)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	// Check if the token is revoked
	if stored.Revoked {
		return nil, ErrTokenRevoked
	}

	// Check if the token is expired
	if time.Now().After(stored.ExpiresAt) {
		return nil, ErrInvalidRefreshToken
	}

	// Get user by ID from claims
	userId := uuid.MustParse(claims.Subject)

	// Get user by ID from claims
	user, err := s.userRepo.FindByID(ctx, userId)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Generate new token pair
	tokenPair, _, err := utils.GenerateTokenPair(user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return nil, err
	}

	// Revoke all previous refresh tokens for the user
	err = s.refreshRepo.RevokeByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	// Create a new refresh token record in the database
	newHash := utils.SHA256Hash(tokenPair.RefreshToken)

	// Create a new refresh token record in the database
	err = s.refreshRepo.Create(ctx, &models.RefreshToken{
		UserID:    user.ID,
		TokenHash: newHash,
		ExpiresAt: time.Now().Add(config.AppConfig.JWTRefreshTTL),
	})

	if err != nil {
		return nil, err
	}

	// Return the new access and refresh tokens
	return &dtoJWT.RefreshResult{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (s *AuthService) GetCurrentUser(ctx context.Context, userId string) (*models.User, error) {
	id := uuid.MustParse(userId)

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
