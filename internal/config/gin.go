package config

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Port         string
}

type Server struct {
	engine *gin.Engine
	config GinConfig
}

// NewServer creates a new HTTP server using Gin.
func NewServer(config GinConfig) *Server {
	engine := gin.New()

	// Middleware: recovery, then logger, then cors
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(cors.Default())

	return &Server{
		engine: engine,
		config: config,
	}
}

// Start starts the HTTP server with the configured timeouts.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.engine,
		ReadTimeout:  s.config.ReadTimeout,
		WriteTimeout: s.config.WriteTimeout,
	}
	return srv.ListenAndServe()
}

// App returns the Gin engine for registering routes and middleware.
func (s *Server) App() *gin.Engine {
	return s.engine
}
