package application

import (
	"codelab/go-gin-template/internal/api/handler"
)

// handler.AuthHandler

type AppHandler struct {
	user handler.UserHandler
}

func SetupHandlers(svc *AppServices) *AppHandler {
	return &AppHandler{
		user: handler.NewUserHandler(svc.user),
	}
}
