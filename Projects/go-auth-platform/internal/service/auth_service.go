package service

import "go-auth-platform/internal/repository"

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

