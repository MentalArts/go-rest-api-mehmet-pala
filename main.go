package main

import (
	"log"
	"os"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/routes"
	"github.com/gin-gonic/gin"
)

// @title Book Library API
// @version 1.0
// @description REST API for managing a book library
// @host localhost:8080
// @BasePath /api/v1
func main() {
	db.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
