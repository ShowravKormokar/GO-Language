package handler

import (
	"encoding/json"
	admDto "go-auth-platform/internal/dto/admin"
	dto "go-auth-platform/internal/dto/common"
	pgDto "go-auth-platform/internal/dto/paginated"
	"go-auth-platform/internal/service"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	admService *service.AdminService
}

func NewAdminHandler(s *service.AdminService) *AdminHandler {
	return &AdminHandler{
		admService: s,
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
