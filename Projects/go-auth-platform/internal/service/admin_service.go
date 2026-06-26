package service

import (
	"context"
	"errors"
	admDto "go-auth-platform/internal/dto/admin"
	pgDto "go-auth-platform/internal/dto/paginated"
	usrDto "go-auth-platform/internal/dto/user"
	"go-auth-platform/internal/mapper"
	"go-auth-platform/internal/repository"
	"math"

	"github.com/google/uuid"
)

type AdminService struct {
	admUsrRepo  repository.AdminUserRepository
	userRepo    repository.UserRepository
	roleRepo    repository.RoleRepository
	refreshRepo repository.RefreshTokenRepository
}

func NewAdminUserService(admUsrRepo repository.AdminUserRepository, userRepo repository.UserRepository, roleRepo repository.RoleRepository, refreshRepo repository.RefreshTokenRepository) *AdminService {
	return &AdminService{
		admUsrRepo:  admUsrRepo,
		userRepo:    userRepo,
		roleRepo:    roleRepo,
		refreshRepo: refreshRepo,
	}
}

func (s *AdminService) GetUsers(ctx context.Context, q admDto.AdminUserQuery) (pgDto.PaginatedResponse[admDto.AdminUserResponse], error) {
	user, total, err := s.admUsrRepo.ListUsers(ctx, q)
	if err != nil {
		return pgDto.PaginatedResponse[admDto.AdminUserResponse]{}, err
	}

	var result []admDto.AdminUserResponse
	for _, u := range user {
		result = append(result, admDto.AdminUserResponse{
			ID:    u.ID.String(),
			Name:  u.Name,
			Email: u.Email,
			Role: usrDto.RoleResponse{
				ID:          u.Role.ID,
				Name:        u.Role.Name,
				Description: u.Role.Description,
			},
			IsActive:  u.IsActive,
			CreatedAt: u.CreatedAt,
		})
	}

	return pgDto.PaginatedResponse[admDto.AdminUserResponse]{
		Data: result,
		Pagination: pgDto.PaginationMeta{
			Page:       q.Page,
			Limit:      q.Limit,
			TotalItems: total,
			TotalPages: int(
				math.Ceil(
					float64(total) /
						float64(q.Limit),
				),
			),
			HasNext: q.Page*q.Limit < int(total),
			HasPrev: q.Page > 1,
		},
	}, nil
}

func (s *AdminService) GetUserByID(ctx context.Context, id string) (*usrDto.UserProfileResponse, error) {
	userID, err := uuid.Parse(id)

	if err != nil {
		return nil, errors.New(
			"invalid user id",
		)
	}

	user, err := s.admUsrRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	response := mapper.ToUserProfileResponse(user)

	return &response, nil
}

// Get all roles
func (s *AdminService) GetAllRole(ctx context.Context) ([]usrDto.RoleResponse, error) {

	roles, err := s.roleRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var response []usrDto.RoleResponse

	for _, role := range roles {
		response = append(
			response,
			usrDto.RoleResponse{
				ID:          role.ID,
				Name:        role.Name,
				Description: role.Description,
			},
		)
	}
	return response, nil
}

// Update user by ID
func (s *AdminService) UpdateUser(ctx context.Context, userID string, req admDto.AdminUpdateUserRequest) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	// Update Name
	if req.Name != nil {
		if *req.Name == "" {
			return errors.New("name cannot be empty")
		}
		user.Name = *req.Name
	}

	// Update Email
	if req.Email != nil {
		if *req.Email == "" {
			return errors.New("email cannot be empty")
		}

		existing, err := s.userRepo.FindByEmail(ctx, *req.Email)

		if err == nil && existing.ID != user.ID {
			return errors.New("email already exists")
		}
		user.Email = *req.Email
	}

	// Update Role
	if req.RoleID != nil {
		user.RoleID = *req.RoleID
	}

	// Update Active status
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	err = s.userRepo.Update(ctx, user)

	if err != nil {
		return err
	}

	// Force logout all devices
	err = s.refreshRepo.RevokeByUserID(ctx, user.ID)

	return err
}
