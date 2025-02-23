package handlers

import (
	"net/http"
	"strconv"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/models"
	"github.com/gin-gonic/gin"
)

// GetReviewsForBook godoc
// @Summary List all reviews for a specific book
// @Tags reviews
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} []models.Review
// @Router /api/v1/books/{id}/reviews [get]
func GetReviewsForBook(c *gin.Context) {
	bookID := c.Param("id")
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
// @Router /api/v1/books/{id}/reviews [post]
func CreateReview(c *gin.Context) {
	bookID := c.Param("id")
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Set the book ID from the URL parameter
	id, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	review.BookID = uint(id)
	result := db.DB.Create(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update an existing review
// @Tags reviews
// @Accept json
// @Produce json
// @Param id path int true "Review ID"
// @Param review body models.Review true "Review data"
// @Success 200 {object} models.Review
// @Router /api/v1/reviews/{id} [put]
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review
	if err := db.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&review)
	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete a review
// @Tags reviews
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/reviews/{id} [delete]
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
