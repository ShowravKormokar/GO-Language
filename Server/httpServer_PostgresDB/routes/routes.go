package routes

import (
	"encoding/json"
	"net/http"
	"server_postgres/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"mssg":   "Server is alive and healthy.",
			"status": http.StatusOK,
		})
	}).Methods("GET")
	r.HandleFunc("/contacts", services.CreateContact).Methods("POST")
	r.HandleFunc("/contacts", services.GetContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", services.GetContact).Methods("GET")
	r.HandleFunc("/contacts/{id}", services.UpdateContact).Methods("PUT")

	return r
}
