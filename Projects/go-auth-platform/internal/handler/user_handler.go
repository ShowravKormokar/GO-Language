package handler

import (
	"go-auth-platform/internal/constants"
	cmmRes "go-auth-platform/internal/dto/common"
	"go-auth-platform/internal/service"
	"go-auth-platform/internal/utils"
	"net/http"
)

type UserHandler struct {
	userService *service.AuthService
}

func NewUserHandler(userService *service.AuthService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Get-Me handler
func (u *UserHandler) Me(rw http.ResponseWriter, rq *http.Request) {
	// Get user id from context
	userID := rq.Context().Value(constants.ContextUserID).(string)

	// Get user by id
	user, err := u.userService.GetCurrentUser(rq.Context(), userID)
	if err != nil {
		http.Error(
			rw,
			"user not found",
			http.StatusNotFound,
		)
		return
	}

	// Success response
	utils.JSON(
		rw,
		http.StatusOK,
		cmmRes.APIResponse[any]{
			Success: true,
			Message: "profile fetched",
			Data:    user,
		},
	)
}
