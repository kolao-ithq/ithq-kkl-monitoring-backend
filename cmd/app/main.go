package main

import (
	"codelab/go-gin-template/cmd/application"
	"codelab/go-gin-template/internal/config"
	"log"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load configuration: %v", err)
	}

	// Initialize HTTP server (Gin)
	ginConfig := config.GinConfig{
		Port:         cfg.Server.Port,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}
	server := config.NewServer(ginConfig)

	repositories := application.SetupRepositories(cfg)
	defer application.CloseRepositories(repositories)
	services := application.SetupServices(repositories)
	handlers := application.SetupHandlers(services)
	application.SetupRoutes(server.App(), handlers)

	// // Close database connection when application exits
	// defer func() {
	// 	if db != nil {
	// 		sqlDB, err := db.DB()
	// 		if err == nil {
	// 			if err := sqlDB.Close(); err != nil {
	// 				log.Printf("Error closing database connection: %v", err)
	// 			} else {
	// 				log.Println("Database connection closed successfully")
	// 			}
	// 		}
	// 	}
	// }()

	// // Start server in a goroutine
	// addr := fmt.Sprintf(":%d", globalEnv.App.Port)
	// go func() {
	// 	log.Printf("Server starting on %s", addr)
	// 	if err := fiber.Listen(addr); err != nil {
	// 		log.Printf("Server error: %v", err)
	// 	}
	// }()

	// // Wait for interrupt signal to gracefully shutdown the server
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server...")

	// // Gracefully shutdown server with a timeout
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// if err := fiber.ShutdownWithContext(ctx); err != nil {
	// 	log.Printf("Server forced to shutdown: %v", err)
	// } else {
	// 	log.Println("Server exited gracefully")
	// }
}
