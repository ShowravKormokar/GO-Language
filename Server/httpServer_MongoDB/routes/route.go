package routes

import (
	"encoding/json"
	"net/http"
	"server_MongoDB/services"

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
	r.HandleFunc("/products", services.CreateProduct).Methods("POST")

	return r
}

/*
curl -X POST http://localhost:8000/products \ 
	-H "Content-Type: application/json" \ 
	-d '{"title":"I Phone","description":"The latest model phone","price":1200.50,"quantity":10,"status":true}'
*/
