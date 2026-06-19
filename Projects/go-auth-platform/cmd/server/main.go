package main

import (
	"fmt"
	"go-auth-platform/internal/config"
	"go-auth-platform/internal/handler"
	"go-auth-platform/internal/migrations"
	"go-auth-platform/internal/repository"
	"go-auth-platform/internal/routes"
	"go-auth-platform/internal/service"
	"net/http"
)

func main() {
	// Load environment variables
	config.LoadEnv()
	// Connect to the database
	config.ConnectDatabase()

	// Run migrations
	if err := migrations.RunMigrations(); err != nil {
		panic(err)
	}
	if err := migrations.SeedRoles(); err != nil {
		panic(err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(config.DB)
	roleRepo := repository.NewRoleRepository(config.DB)
	refreshRepo := repository.NewRefreshTokenRepository(config.DB)
	blacklistRepo := repository.NewBlacklistRepository(config.DB)

	// Initialize services
	authService := service.NewAuthService(userRepo, roleRepo, refreshRepo, blacklistRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(authService)

	// Register routes
	r := routes.RegisterRouter(authHandler, userHandler, blacklistRepo)

	fmt.Printf("%s running on port %s\n", config.AppConfig.AppName, config.AppConfig.AppPort)
	if err := http.ListenAndServe(":"+config.AppConfig.AppPort, r); err != nil {
		fmt.Println("Server couldn't connected!", err)
	}
}
