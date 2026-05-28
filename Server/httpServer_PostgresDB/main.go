package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server_postgres/database"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"mssg":   "Server is alive and healthy.",
			"status": http.StatusOK,
		})
	})

	// Database Connect
	database.ConnectPostgres()

	fmt.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Couldn't connect to server:", err)
	}
}
