package config

// @title ITQ HR Landing Page Backend API
// @version 1.0.0
// @description Backend API for ITQ HR Landing Page application with authentication, user management, and candidate management features

// IMPORTANT: All API endpoints require an Access token for authentication.
// Please contact the administrator to obtain your API key.
// Include the API key in the X-API-Key header for all requests.

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

import (
	"fmt"
	// _ "codelab/go-gin-template/pkg/swagger/docs"
)

// SwaggerConfig holds Swagger configuration
type SwaggerConfig struct {
	Title       string
	Description string
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
}

// NewSwaggerConfig creates a new Swagger configuration
func NewSwaggerConfig(serverConfig ServerConfig) *SwaggerConfig {
	var host string
	switch serverConfig.Env {
	case "prd":
		host = fmt.Sprintf("localhost:%s", serverConfig.Port)
	case "dev":
		host = fmt.Sprintf("localhost:%s", serverConfig.Port)
	case "uat":
		host = fmt.Sprintf("localhost:%s", serverConfig.Port)
	default:
		host = fmt.Sprintf("localhost:%s", serverConfig.Port)
	}
	return &SwaggerConfig{
		Title:       "ITQ HR Landing Page Backend API",
		Description: "Backend API for ITQ HR Landing Page application",
		Version:     "1.0.0",
		Host:        host,
		BasePath:    "/api/v1",
		Schemes:     []string{"http", "https"},
	}
}
