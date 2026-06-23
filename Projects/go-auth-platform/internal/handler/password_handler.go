package handler

import (
	"encoding/json"
	dto "go-auth-platform/internal/dto/common"
	dtoPR "go-auth-platform/internal/dto/user"
	"go-auth-platform/internal/service"
	"go-auth-platform/internal/utils"
	"net/http"
)

// Dependency Inject
type PasswordHandler struct {
	passwordService service.PasswordService
}

// Constructor
func NewPasswordHandler(passwordService service.PasswordService) *PasswordHandler {

	return &PasswordHandler{
		passwordService: passwordService,
	}
}

// Forgot password handlers
func (h *PasswordHandler) ForgotPassword(rw http.ResponseWriter, rq *http.Request) {
	var req dtoPR.ForgotPasswordRequest

	if err := json.NewDecoder(rq.Body).Decode(&req); err != nil {
		http.Error(
			rw,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	link, err := h.passwordService.ForgotPassword(rq.Context(), req.Email)
	if err != nil {
		utils.JSON(
			rw,
			http.StatusInternalServerError,
			err,
		)
	}
	// In production the reset link sent on user email
	utils.JSON(
		rw,
		http.StatusOK,
		dto.APIResponse[string]{
			Success: true,
			Message: "If email exists reset link sent",
			Data:    link, // This is only for learning purpose
		},
	)
}
