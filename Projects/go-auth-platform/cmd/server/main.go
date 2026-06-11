package main

import (
	"fmt"
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/migrations"
	"go-auth-platform/internal/routes"
	"net/http"
)

func main() {
	// Load environment variables
	config.LoadEnv()
	// Connect to the database
	config.ConnectDatabase()

	err := migrations.RunMigrations()
	if err != nil {
		panic(err)
	}

	err = migrations.SeedRoles()
	if err != nil {
		panic(err)
	}

	// Register routes and start the server
	r := routes.RegisterRouter()

	fmt.Printf("%s running on port %s\n", config.AppConfig.AppName, config.AppConfig.AppPort)
	err = http.ListenAndServe(":"+config.AppConfig.AppPort, r)
	if err != nil {
		fmt.Println("Server couldn't connected!", err)
	}
}
