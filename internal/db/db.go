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

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatal("Database connection environment variables are not set correctly.")
	}

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

	err = DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	addForeignKey("books", "author_id", "authors(id)", "CASCADE")
	addForeignKey("reviews", "book_id", "books(id)", "CASCADE")
}

func addForeignKey(table, field, ref, onDelete string) {
	var count int64
	err := DB.Raw(fmt.Sprintf(`
        SELECT COUNT(*)
        FROM information_schema.table_constraints
        WHERE constraint_type = 'FOREIGN KEY'
        AND table_name = '%s'
        AND constraint_name = 'fk_%s_%s'`,
		table, table, field)).Scan(&count).Error

	if err != nil {
		log.Printf("Warning: Could not check foreign key constraint for %s: %v", table, err)
		return
	}

	if count == 0 {
		err = DB.Exec(fmt.Sprintf(`
            ALTER TABLE %s 
            ADD CONSTRAINT fk_%s_%s 
            FOREIGN KEY (%s) 
            REFERENCES %s 
            ON DELETE %s`,
			table, table, field, field, ref, onDelete)).Error

		if err != nil {
			log.Printf("Warning: Could not add foreign key constraint to %s: %v", table, err)
		}
	} else {
		log.Printf("Foreign key constraint fk_%s_%s already exists on table %s", table, field, table)
	}
}
