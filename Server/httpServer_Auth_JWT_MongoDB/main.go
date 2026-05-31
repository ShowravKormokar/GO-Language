package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 1. Connect to MongoDB
	// 2. Set up HTTP server with routes for registration and login
	// 3. Implement JWT authentication for protected routes

	fmt.Println("Server running on port:8000")
	err := http.ListenAndServe(os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println("Server couldn't connect:", err)
	}
}
