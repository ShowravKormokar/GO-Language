package routes

import (
	"encoding/json"
	"httpServer_JWT_MongoDB/dto"
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
	r.HandleFunc("/auth/register", services.RegisterAuthService).Methods("POST")
	r.HandleFunc("/auth/login", services.LoginAuthService).Methods("POST")

	return r
}
