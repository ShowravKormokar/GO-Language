package main

import (
	"fmt"
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/routes"
	"net/http"
)

func main() {
	// Load environment variables
	config.LoadEnv()
	// Connect to the database
	config.ConnectDatabase()
	// Register routes and start the server
	r := routes.RegisterRouter()

	fmt.Println("Server connected successfully on port:", config.AppConfig.AppPort)
	err := http.ListenAndServe(":"+config.AppConfig.AppPort, r)
	if err != nil {
		fmt.Println("Server couldn't connected!", err)
	}
}
