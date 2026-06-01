package main

import (
	"fmt"
	"httpServer_JWT_MongoDB/routes"
	"net/http"
)

func main() {
	// 1. Connect to MongoDB
	// 2. Set up HTTP server with routes for registration and login
	// 3. Implement JWT authentication for protected routes
	r := routes.RegisterRoutes()

	fmt.Println("Server running on port:", 8000)
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Server couldn't connect:", err)
	}
}
