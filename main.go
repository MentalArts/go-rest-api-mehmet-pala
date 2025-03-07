package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings" // Added for path checking
	"time"

	_ "github.com/MentalArts/go-rest-api-mehmet-pala/docs" // Import your generated docs package
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
		Addr:     "redis:6379", // Docker'daki Redis servisi
		Password: "",           // Varsayılan şifre boş
		DB:       0,            // Varsayılan veritabanı
	})
}

// Rate limiting middleware
func rateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip rate-limiting for Swagger routes
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			c.Next()
			return
		}

		ip := c.ClientIP() // Kullanıcının IP adresini al
		ctx := context.Background()

		// Redis'teki anahtar
		key := fmt.Sprintf("rate_limit:%s", ip)

		// Redis'ten mevcut istek sayısını al
		count, err := rdb.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			log.Printf("Error reading from Redis: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		// Eğer istek sayısı limitin üzerinde ise 429 Too Many Requests döner
		if count >= limit {
			c.JSON(429, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Eğer limit aşılmadıysa, isteği say
		_, err = rdb.Set(ctx, key, count+1, window).Result()
		if err != nil {
			log.Printf("Error setting value in Redis: %v", err)
			c.JSON(500, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		// İstek sayısını artırarak işlemi devam ettir
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
