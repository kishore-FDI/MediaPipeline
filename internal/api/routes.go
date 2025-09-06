package api

import (
	"net/http"
	"time"

	"mediapipeline/internal/config"
	"mediapipeline/internal/db"
	"mediapipeline/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	r.Use(corsMiddleware())
	r.GET("/health", healthCheck)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Media Pipeline API",
				"version": "1.0.0",
				"service": "Content Moderation Pipeline",
			})
		})

		uploads := v1.Group("/uploads")
		uploads.Use(middleware.RateLimiter(db.RDB, 10, time.Minute, middleware.BusinessRateLimit{}))
		{
			uploads.POST("/", uploadHandler)
			uploads.PUT("/:id", resumeUploadHandler)
			uploads.GET("/:id/status", statusHandler)
		}

		storage := v1.Group("/storage")
		{
			storage.GET("/:id", downloadHandler)
			storage.DELETE("/:id", deleteHandler)
		}

		moderation := v1.Group("/moderation")
		{
			moderation.POST("/check", moderationHandler)
			moderation.GET("/:id/result", resultHandler)
		}

		SetupBusinessRoutes(v1)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "Media Pipeline API",
	})
}

func uploadHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Upload handler not implemented yet"})
}

func resumeUploadHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Resume upload handler not implemented yet"})
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Status handler not implemented yet"})
}

func downloadHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Download handler not implemented yet"})
}

func deleteHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Delete handler not implemented yet"})
}

func moderationHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Moderation handler not implemented yet"})
}

func resultHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Result handler not implemented yet"})
}
