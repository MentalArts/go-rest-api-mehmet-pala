package routes

import (
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		// Book endpoints
		api.GET("/books", handlers.GetBooks)
		api.GET("/books/:id", handlers.GetBookByID)
		api.POST("/books", handlers.CreateBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)

		// Author endpoints
		api.GET("/authors", handlers.GetAuthors)
		api.GET("/authors/:id", handlers.GetAuthorByID)
		api.POST("/authors", handlers.CreateAuthor)
		api.PUT("/authors/:id", handlers.UpdateAuthor)
		api.DELETE("/authors/:id", handlers.DeleteAuthor)

		// Review endpoints
		api.GET("/books/:id/reviews", handlers.GetReviewsForBook)
		api.POST("/books/:id/reviews", handlers.CreateReview)
		api.PUT("/reviews/:id", handlers.UpdateReview)
		api.DELETE("/reviews/:id", handlers.DeleteReview)
	}

}
