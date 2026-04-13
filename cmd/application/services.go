package application

import (
	"codelab/go-gin-template/internal/core/modules/users/ports/input"
	"codelab/go-gin-template/internal/core/modules/users/services"
)

// AppServices holds all application services
type AppServices struct {
	user input.UserService
	// auth input.AuthService // Use concrete type for direct navigation
}

// SetupServices initializes all core business logic services
func SetupServices(repos *AppRepositories) *AppServices {
	userService := services.NewUserService(repos.users, repos.redis)
	// authService := services.NewAuthService(repos.user, repos.redis)
	return &AppServices{
		user: userService,
		// auth: authService,
	}
}
