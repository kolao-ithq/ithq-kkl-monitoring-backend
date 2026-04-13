package router

import (
	"codelab/go-gin-template/internal/api/handler"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes registers user routes on the given Gin router (engine or group).
func SetupUserRoutes(r gin.IRouter, userHandler handler.UserHandler) {
	r.GET("/users", wrap(userHandler.GetUser))
	r.GET("/users/:id", wrap(userHandler.GetUserByID))
	r.PUT("/users", wrap(userHandler.UpdateUser))
	r.DELETE("/users/:id", wrap(userHandler.DeleteUser))
}

// wrap adapts a handler that returns error to gin.HandlerFunc.
func wrap(fn func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = fn(c)
	}
}
