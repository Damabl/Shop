package routes

import (
	"Shop/internal/auth"
	"Shop/internal/handlers"
	"Shop/internal/repositories"
	"Shop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoutes(router *gin.Engine, db *gorm.DB) {
	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	handler := handlers.NewProductHandler(service)
	routes := router.Group("/products")
	{
		routes.Use(auth.AuthMiddleware())
		routes.POST("/product", auth.RoleMiddleware("Admin"), handler.CreateProduct)
		routes.PUT("/product/:id", auth.RoleMiddleware("Admin"), handler.UpdateProduct)
		routes.DELETE("/product/:id", auth.RoleMiddleware("Admin"), handler.DeleteProduct)
		routes.GET("/product/:id", handler.GetProductByID)
		routes.GET("", handler.GetAllProducts)
	}
}
