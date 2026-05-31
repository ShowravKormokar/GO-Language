package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Connect to MongoDB
	// 2. Set up HTTP server with routes for registration and login
	// 3. Implement JWT authentication for protected routes

	fmt.Println("Server running on port:8080")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server couldn't connect:", err)
	}
}
