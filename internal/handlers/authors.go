package handlers

import (
	"net/http"
	"strconv"

	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/db"
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/models"
	"github.com/gin-gonic/gin"
)

// GetAuthors godoc
// @Summary List all authors
// @Tags authors
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/authors [get]
func GetAuthors(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	var authors []models.Author
	result := db.DB.Offset(offset).Limit(limit).Find(&authors)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": authors, "page": page, "limit": limit})
}

// GetAuthorByID godoc
// @Summary Get a single author by ID
// @Tags authors
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Router /api/v1/authors/{id} [get]
func GetAuthorByID(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	result := db.DB.First(&author, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": author})
}

// CreateAuthor godoc
// @Summary Create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body models.Author true "Author to create"
// @Success 201 {object} models.Author
// @Router /api/v1/authors [post]
func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.DB.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": author})
}

// UpdateAuthor godoc
// @Summary Update an existing author
// @Tags authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body models.Author true "Author data"
// @Success 200 {object} models.Author
// @Router /api/v1/authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := db.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Save(&author)
	c.JSON(http.StatusOK, gin.H{"data": author})
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Tags authors
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} map[string]string
// @Router /api/v1/authors/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author
	if err := db.DB.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}
	db.DB.Delete(&author)
	c.JSON(http.StatusOK, gin.H{"data": "Author deleted"})
}
