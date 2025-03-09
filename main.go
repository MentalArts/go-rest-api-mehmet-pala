package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/MentalArts/go-rest-api-mehmet-pala/docs"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Prometheus metrics
var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_requests_total",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint"},
	)
)

// Redis client
var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	pong, err := rdb.Ping(context.Background()).Result()
	log.Printf("Redis ping: %s, error: %v", pong, err)
}

// Rate limiting middleware
func rateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip rate-limiting for Swagger routes
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.Next()
			return
		}

		ip := c.ClientIP() // Get client IP
		ctx := context.Background()
		key := fmt.Sprintf("rate_limit:%s", ip)

		// Increment count atomically and set expiration if new key
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			log.Printf("Error incrementing Redis count: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		// If this is the first request, set expiration
		if count == 1 {
			_, err = rdb.Expire(ctx, key, window).Result()
			if err != nil {
				log.Printf("Error setting Redis expiration: %v", err)
				c.JSON(500, gin.H{"error": "Internal server error"})
				c.Abort()
				return
			}
		}

		// Check if limit exceeded
		if count > int64(limit) {
			c.JSON(429, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Proceed with request
		c.Next()
	}
}

func setCache(key string, value string) {
	err := rdb.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		log.Fatalf("Redis error: %v", err)
	}
}

func getCache(key string) (string, error) {
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// @title Book Library API
// @version 1.0
// @description REST API for managing a book library
// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
func main() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Initialize Redis
	initRedis()

	// Initialize database connection
	db.InitDB()

	// Setup Prometheus metrics
	prometheus.MustRegister(requestCount)

	// Initialize the router
	r := gin.Default()

	// Rate limiting middleware: 5 istek / 1 dakika
	r.Use(rateLimitMiddleware(5, time.Minute))

	// Set trusted proxies
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Serve Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Expose Prometheus metrics using gin's WrapH method
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Setup routes
	routes.SetupRoutes(r)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

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
