package http

import (
	"context"
	_ "github.com/dan-ibm/go-crud-app/docs"
	"github.com/dan-ibm/go-crud-app/internal/domain"
	"github.com/dan-ibm/go-crud-app/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Books interface {
	Create(ctx context.Context, book domain.BookInput) error
	GetByID(ctx context.Context, id int64) (domain.Book, error)
	GetAll(ctx context.Context) ([]domain.Book, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.BookInput) error
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getAllBooks)
			books.GET("/user", h.getAllBooksByUser)
			books.GET("/:id", h.getBookById)
			books.PUT("/:id", h.updateBook)
			books.DELETE("/:id", h.deleteBook)
		}
	}

	return router
}
