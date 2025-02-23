package main

import (
	"log"
	"os"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Book Library API
// @version 1.0
// @description REST API for managing a book library
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize database connection
	db.InitDB()

	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Starting server on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
