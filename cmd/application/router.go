package application

import (
	"codelab/go-gin-template/internal/api/router"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine, handlers *AppHandler) {
	engine.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World! 👋")
	})
	// API group
	api := engine.Group("/api/v1")
	router.SetupUserRoutes(api, handlers.user)
}
