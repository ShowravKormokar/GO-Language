package routes

import (
	"encoding/json"
	"httpServer_JWT_MongoDB/dto"
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

	return r
}
