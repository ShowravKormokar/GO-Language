package handler

import (
	"encoding/json"
	admDto "go-auth-platform/internal/dto/admin"
	dto "go-auth-platform/internal/dto/common"
	pgDto "go-auth-platform/internal/dto/paginated"
	urdto "go-auth-platform/internal/dto/user"

	"go-auth-platform/internal/service"
	"go-auth-platform/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	admService  *service.AdminService
	userService *service.UserService
}

func NewAdminHandler(s *service.AdminService, u *service.UserService) *AdminHandler {
	return &AdminHandler{
		admService:  s,
		userService: u,
	}
}

// Get all users for admin only
func (h *AdminHandler) GetUsers(rw http.ResponseWriter, rq *http.Request) {
	query := admDto.AdminUserQuery{
		Page:   1,
		Limit:  20,
		Search: rq.URL.Query().Get("search"),
		Role:   rq.URL.Query().Get("role"),
		Sort:   rq.URL.Query().Get("sort"),
		Order:  rq.URL.Query().Get("order"),
	}

	if p := rq.URL.Query().Get("page"); p != "" {
		query.Page, _ = strconv.Atoi(p)
	}

	if l := rq.URL.Query().Get("limit"); l != "" {
		query.Limit, _ = strconv.Atoi(l)
	}

	if a := rq.URL.Query().Get("is_active"); a != "" {
		v, _ := strconv.ParseBool(a)
		query.IsActive = &v
	}

	result, err := h.admService.GetUsers(rq.Context(), query)
	if err != nil {
		http.Error(
			rw,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	rw.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(rw).Encode(
		dto.APIResponse[pgDto.PaginatedResponse[admDto.AdminUserResponse]]{
			Success: true,
			Message: "users fetched successfully",
			Data:    result,
		},
	)
}

// Get specific user by id
func (h *AdminHandler) GetUserById(rw http.ResponseWriter, rq *http.Request) {
	id := mux.Vars(rq)["id"]

	user, err := h.admService.GetUserByID(rq.Context(), id)
	if err != nil {
		utils.JSON(
			rw,
			http.StatusNotFound,
			dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	utils.JSON(
		rw,
		http.StatusOK,
		dto.APIResponse[*urdto.UserProfileResponse]{
			Success: true,
			Message: "user fetched successfully",
			Data:    user,
		},
	)

}

// Delete user by ID (Soft Delete)
func (h *AdminHandler) DeleteUser(rw http.ResponseWriter, rq *http.Request) {

	id := mux.Vars(rq)["id"]

	err := h.userService.DeleteUser(rq.Context(), id)

	if err != nil {
		utils.JSON(
			rw,
			http.StatusBadRequest,
			dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	rw.WriteHeader(
		http.StatusNoContent,
	)
}

// Get all role
func (h *AdminHandler) GetRoles(rw http.ResponseWriter, rq *http.Request) {
	roles, err := h.admService.GetAllRole(rq.Context())
	if err != nil {
		utils.JSON(
			rw,
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Success: false,
				Message: "failed to fetch roles",
			},
		)
		return
	}

	utils.JSON(
		rw,
		http.StatusOK,
		dto.APIResponse[[]urdto.RoleResponse]{
			Success: true,
			Message: "roles fetched successfully",
			Data:    roles,
		},
	)
}
