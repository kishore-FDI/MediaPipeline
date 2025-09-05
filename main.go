package main

import (
	"log"
	"os"
	"time"

	"mediapipeline/internal/api"
	"mediapipeline/internal/config"
	"mediapipeline/internal/db"
	"mediapipeline/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Init Redis
	db.InitRedis()

	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Apply rate limiter (example: 100 requests per minute per IP)
	r.Use(middleware.RateLimiter(db.RDB, 1, time.Minute))

	api.SetupRoutes(r, cfg)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting Media Pipeline API server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
