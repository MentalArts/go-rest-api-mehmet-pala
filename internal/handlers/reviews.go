package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetReviewsForBook godoc
// @Summary List all reviews for a specific book
// @Tags reviews
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} []models.Review
// @Router /books/{id}/reviews [get]
func GetReviewsForBook(c *gin.Context) {
	bookID := c.Param("id")

	// Check if book exists before fetching reviews
	var book models.Book
	if err := db.DB.First(&book, "id = ?", bookID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var reviews []models.Review
	result := db.DB.Where("book_id = ?", bookID).Find(&reviews)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// CreateReview godoc
// @Summary Create a new review for a book
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param review body models.Review true "Review to create"
// @Success 201 {object} models.Review
// @Router /books/{id}/reviews [post]
func CreateReview(c *gin.Context) {
	bookID := c.Param("id")
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate Book ID
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	// Check if book exists
	var book models.Book
	if err := db.DB.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	review.BookID = uint(id)

	// Create review
	result := db.DB.Create(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": review})
}

func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	// Fetch review
	if err := db.DB.First(&review, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Bind JSON request
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save updated review
	if err := db.DB.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete a review
// @Tags reviews
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} map[string]string
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := db.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	db.DB.Delete(&review)
	c.JSON(http.StatusOK, gin.H{"data": "Review deleted"})
}
