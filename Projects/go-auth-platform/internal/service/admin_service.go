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
	admUsrRepo repository.AdminUserRepository
}

func NewAdminUserService(admUsrRepo repository.AdminUserRepository) *AdminService {
	return &AdminService{
		admUsrRepo: admUsrRepo,
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
