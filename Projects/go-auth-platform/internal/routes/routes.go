package routes

import (
	"encoding/json"
	"go-auth-platform/internal/handler"
	"go-auth-platform/internal/middleware"
	"go-auth-platform/internal/repository"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

var serverStartTime = time.Now()

func healthHandler(rw http.ResponseWriter, rq *http.Request) {
	start := time.Now()

	// Memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	cpuCount := runtime.NumCPU()
	goroutines := runtime.NumGoroutine()

	// Version info
	version := "1.0.0v"

	response := map[string]interface{}{
		"success":    true,
		"message":    "Server is healthy and alive",
		"speed_ms":   time.Since(start).Milliseconds(),
		"uptime":     time.Since(serverStartTime).String(),
		"time":       time.Now().Format(time.RFC3339),
		"memory_mb":  m.Alloc / 1024 / 1024,
		"cpu_count":  cpuCount,
		"goroutines": goroutines,
		"version":    version,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

func RegisterRouter(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	passwordHandler *handler.PasswordHandler,

	blacklistRepo repository.BlacklistRepository,
) *mux.Router {
	r := mux.NewRouter()

	// App health check
	r.HandleFunc("/health", healthHandler).Methods("GET")

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

	return r
}
