package routes

import (
	"go-auth-platform/internal/handler"
	"go-auth-platform/internal/middleware"
	"go-auth-platform/internal/repository"

	"github.com/gorilla/mux"
)

func RegisterRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	passwordHandler *handler.PasswordHandler,

	blacklistRepo repository.BlacklistRepository,
) *mux.Router {
	r := mux.NewRouter()

	// App health check
	r.HandleFunc("/health", handler.ServerHealthHandler).Methods("GET")

	// Auth routes
	auth := r.PathPrefix("/api/v1/auth").Subrouter()

	// Public routes
	// User Register Route
	auth.HandleFunc("/register", authHandler.Register).Methods("POST")
	// User Login Route
	auth.HandleFunc("/login", authHandler.Login).Methods("POST")
	// Token Refresh Route
	auth.HandleFunc("/refresh", authHandler.Refresh).Methods("POST")
	// Forgot Password Route
	auth.HandleFunc("/forgot-password", passwordHandler.ForgotPassword).Methods("POST")
	// Reset Password Route
	auth.HandleFunc("/reset-password", passwordHandler.ResetPassword).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/api/v1/").Subrouter()
	// The AuthRequired middleware is applied to the protected subrouter
	protected.Use(middleware.AuthRequired(blacklistRepo))

	// User Logout Route
	protected.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	// Get Current User Route
	protected.HandleFunc("/users/me", userHandler.Me).Methods("GET")
	// Change password - Current login user
	protected.HandleFunc("/users/me/password", userHandler.ChangePassword).Methods("PATCH")

	// Admin Routes [Protected: Auth + Role]
	admin := r.PathPrefix("/api/v1/admin").Subrouter()
	admin.Use(middleware.AuthRequired(blacklistRepo)) // Auth middleware
	admin.Use(middleware.RequireMinRole("admin"))     // Role middleware

	// Get all users route
	admin.HandleFunc("/users", adminHandler.GetUsers).Methods("GET")

	return r
}
