package handler

import (
	"encoding/json"
	"go-auth-platform/internal/constants"
	dtoJWT "go-auth-platform/internal/dto/claims"
	cmmRes "go-auth-platform/internal/dto/common"
	dto "go-auth-platform/internal/dto/common"
	urdto "go-auth-platform/internal/dto/user"
	"go-auth-platform/internal/service"
	"go-auth-platform/internal/utils"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
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

// Change password handler
func (u *UserHandler) ChangePassword(rw http.ResponseWriter, rq *http.Request) {
	var req urdto.ChangePasswordRequest

	// Decode to get user req (request body)
	err := json.NewDecoder(rq.Body).Decode(&req)
	if err != nil {

		utils.JSON(
			rw,
			400,
			dto.ErrorResponse{
				Success: false,
				Message: "invalid body",
			},
		)
		return
	}

	// Get current logged in user-id
	userID, ok := rq.Context().Value(constants.ContextUserID).(string)
	if !ok {
		utils.JSON(
			rw,
			http.StatusUnauthorized,
			dto.ErrorResponse{
				Success: false,
				Message: "unauthorized",
			},
		)
		return
	}

	// claims
	claims, ok := rq.Context().Value(constants.ContextClaims).(*dtoJWT.JWTClaims)
	if !ok {
		utils.JSON(
			rw,
			http.StatusUnauthorized,
			dto.ErrorResponse{
				Success: false,
				Message: "invalid token claims",
			},
		)
		return
	}

	// Change password using service method
	err = u.userService.ChangePassword(rq.Context(), userID, claims, req)
	if err != nil {
		utils.JSON(
			rw,
			http.StatusInternalServerError,
			dto.ErrorResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	// Success response
	utils.JSON(
		rw,
		http.StatusOK,
		dto.APIResponse[any]{
			Success: true,
			Message: "password changed successfully, please login",
		},
	)
}
