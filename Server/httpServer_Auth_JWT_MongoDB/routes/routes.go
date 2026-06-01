package routes

import (
	"encoding/json"
	"httpServer_JWT_MongoDB/dto"
	"httpServer_JWT_MongoDB/middleware"
	"httpServer_JWT_MongoDB/services"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/health", func(rw http.ResponseWriter, rq *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(dto.BasicResponse{
			Success: true,
			Message: "Server is alive and healthy",
		})
	}).Methods("GET")
	r.HandleFunc("/api/auth/register", services.RegisterAuthService).Methods("POST")
	r.HandleFunc("/api/auth/login", services.LoginAuthService).Methods("POST")

	// Protected routes
	r.Handle("/api/users/profile", middleware.JWTMiddleware(http.HandlerFunc(services.UserProfile))).Methods("GET")
	r.Handle("/api/admin/users", middleware.JWTMiddleware(http.HandlerFunc(services.GetAllUsers))).Methods("GET")

	return r
}
