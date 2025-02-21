package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	maxRetries := 5
	var err error

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Database connection attempt %d/%d failed: %v", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})

	// Add foreign key constraints
	addForeignKey("books", "author_id", "authors(id)", "CASCADE")
	addForeignKey("reviews", "book_id", "books(id)", "CASCADE")
}

func addForeignKey(table, field, ref, onDelete string) {
	err := DB.Exec(fmt.Sprintf(`
		ALTER TABLE %s 
		ADD CONSTRAINT fk_%s_%s 
		FOREIGN KEY (%s) 
		REFERENCES %s 
		ON DELETE %s`,
		table, table, field, field, ref, onDelete)).Error

	if err != nil {
		log.Printf("Warning: Could not add foreign key constraint to %s: %v", table, err)
	}
}
