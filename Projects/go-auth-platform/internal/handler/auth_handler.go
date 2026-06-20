package handler

import (
	"encoding/json"
	"go-auth-platform/internal/constants"
	dto "go-auth-platform/internal/dto/auth"
	claims "go-auth-platform/internal/dto/claims"
	cmmRes "go-auth-platform/internal/dto/common"
	"go-auth-platform/internal/service"
	"go-auth-platform/internal/utils"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register Handler
func (h *AuthHandler) Register(rw http.ResponseWriter, rq *http.Request) {
	var req dto.RegisterRequest

	// Decode req body and check it correct or not
	if err := json.NewDecoder(rq.Body).Decode(&req); err != nil {
		utils.JSON(
			rw,
			http.StatusBadRequest,
			cmmRes.ErrorResponse{
				Success: true,
				Message: "invalid request body",
			},
		)
		return
	}

	// Set user through service
	_, err := h.authService.Register(rq.Context(), req)
	if err != nil {
		utils.JSON(
			rw,
			http.StatusBadRequest,
			cmmRes.ErrorResponse{
				Success: false,
				Message: err.Error(),
			},
		)
		return
	}

	// Success to response as success
	utils.JSON(
		rw,
		http.StatusCreated,
		cmmRes.APIResponse[any]{
			Success: true,
			Message: "user created successfully, please login to continue",
		},
	)
}

// Login Handler
func (h *AuthHandler) Login(rw http.ResponseWriter, rq *http.Request) {
	var req dto.LoginRequest

	// Decode body and check valid or not
	if err := json.NewDecoder(rq.Body).Decode(&req); err != nil {
		http.Error(
			rw,
			"invalid body",
			http.StatusBadRequest,
		)
		return
	}

	// Get successfully logged in user
	result, err := h.authService.Login(rq.Context(), req)
	if err != nil {
		http.Error(
			rw,
			err.Error(),
			http.StatusUnauthorized,
		)
		return
	}

	// Set tokens on cookie
	utils.SetAccessCookie(rw, result.AccessToken)
	utils.SetRefreshCookie(rw, result.RefreshToken)

	// Send success response
	utils.JSON(
		rw,
		http.StatusOK,
		cmmRes.APIResponse[any]{
			Success: true,
			Message: "login success",
			Data:    result.User,
		},
	)
}

// Logout Handler
func (h *AuthHandler) Logout(rw http.ResponseWriter, rq *http.Request) {

	// Get claims
	claims := rq.Context().Value(constants.ContextClaims).(*claims.JWTClaims)

	// Perform logout operation
	err := h.authService.Logout(rq.Context(), claims)
	if err != nil {
		http.Error(
			rw,
			"logout failed",
			http.StatusInternalServerError,
		)
		return
	}

	// Remove tokens form cookie
	utils.ClearAuthCookies(rw)

	// Success response
	utils.JSON(
		rw,
		http.StatusOK,
		cmmRes.APIResponse[any]{
			Success: true,
			Message: "logout success",
		},
	)
}

// Refresh Token Handler
func (h *AuthHandler) Refresh(rw http.ResponseWriter, rq *http.Request) {

	// Get current refresh token
	cookie, err := rq.Cookie("refresh_token")
	if err != nil {
		http.Error(
			rw,
			"missing refresh token",
			http.StatusUnauthorized,
		)
		return
	}

	// Refresh token to get new tokens
	result, err := h.authService.Refresh(rq.Context(), cookie.Value)
	if err != nil {
		http.Error(
			rw,
			"invalid refresh token",
			http.StatusUnauthorized,
		)
		return
	}

	// Set new tokens on cookie
	utils.SetAccessCookie(rw, result.AccessToken)
	utils.SetRefreshCookie(rw, result.RefreshToken)

	// Success response
	utils.JSON(
		rw,
		http.StatusOK,
		cmmRes.APIResponse[any]{
			Success: true,
			Message: "token refreshed",
		},
	)
}
