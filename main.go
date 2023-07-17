package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Database connection setup
	db = setupDatabase()
	defer db.Close()

	// Create the "Feed configuration" table if it doesn't exist
	createTable()

	// Initialize the Gin router
	router := gin.Default()

	router.Use(corsMiddleware())

	// Define API routes
	router.GET("/feed-configurations", getAllFeedConfigurations)
	router.GET("/feed-configurations/:id", getFeedConfiguration)
	router.POST("/feed-configurations", createFeedConfiguration)
	router.PUT("/feed-configurations/:id", updateFeedConfiguration)
	router.DELETE("/feed-configurations/:id", deleteFeedConfiguration)

	// Start the server
	log.Fatal(router.Run(":8000"))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
