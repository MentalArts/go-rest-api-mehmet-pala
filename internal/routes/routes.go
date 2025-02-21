package routes

import (
	"github.com/MentalArts/go-rest-api-mehmet-pala/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/books", handlers.GetBooks)
		api.GET("/books/:id", handlers.GetBookByID)
		api.POST("/books", handlers.CreateBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)

		api.GET("/authors", handlers.GetAuthors)
		api.GET("/authors/:id", handlers.GetAuthorByID)
		api.POST("/authors", handlers.CreateAuthor)
		api.PUT("/authors/:id", handlers.UpdateAuthor)
		api.DELETE("/authors/:id", handlers.DeleteAuthor)

		api.GET("/books/:id/reviews", handlers.GetReviewsForBook)
		api.POST("/books/:id/reviews", handlers.CreateReview)
		api.PUT("/reviews/:id", handlers.UpdateReview)
		api.DELETE("/reviews/:id", handlers.DeleteReview)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
