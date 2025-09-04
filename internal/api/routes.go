package api

import (
	"net/http"

	"mediapipeline/internal/config"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	// Add CORS middleware
	r.Use(corsMiddleware())

	// Health check endpoint
	r.GET("/health", healthCheck)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Basic info endpoint
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Media Pipeline API",
				"version": "1.0.0",
				"service": "Content Moderation Pipeline",
			})
		})

		// Upload routes (to be implemented)
		uploads := v1.Group("/uploads")
		{
			uploads.POST("/", uploadHandler)           // Start upload
			uploads.PUT("/:id", resumeUploadHandler)   // Resume upload
			uploads.GET("/:id/status", statusHandler)  // Get upload status
		}

		// Storage routes (to be implemented)
		storage := v1.Group("/storage")
		{
			storage.GET("/:id", downloadHandler)       // Download file
			storage.DELETE("/:id", deleteHandler)      // Delete file
		}

		// Moderation routes (to be implemented)
		moderation := v1.Group("/moderation")
		{
			moderation.POST("/check", moderationHandler) // Check content
			moderation.GET("/:id/result", resultHandler) // Get moderation result
		}
	}
}

// corsMiddleware adds CORS headers
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

// healthCheck returns the health status of the service
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"service": "Media Pipeline API",
	})
}

// Placeholder handlers (to be implemented in separate files)
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
