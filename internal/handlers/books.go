package handlers

import (
	"net/http"
	"strconv"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/models"
	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary List all books
// @Tags books
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/books [get]
func GetBooks(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	var books []models.Book
	result := db.DB.Offset(offset).Limit(limit).Preload("Author").Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books, "page": page, "limit": limit})
}

// GetBookByID godoc
// @Summary Get a single book by ID
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /api/v1/books/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	result := db.DB.Preload("Author").First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBook godoc
// @Summary Create a new book
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book to create"
// @Success 201 {object} models.Book
// @Router /api/v1/books [post]
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update an existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book data"
// @Success 200 {object} models.Book
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete a book
// @Tags books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	db.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": "Book deleted"})
}
