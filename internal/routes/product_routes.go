package routes

import (
	"Shop/internal/auth"
	"Shop/internal/cloud"
	"Shop/internal/handlers"
	"Shop/internal/repositories"
	"Shop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func SetupProductRoutes(router *gin.Engine, db *gorm.DB) {
	cloudinaryService, err := cloud.NewCloudinaryService()
	repo := repositories.NewProductRepository(db)
	imageService := services.NewImageService(cloudinaryService)
	service := services.NewProductService(repo, imageService)
	if err != nil {
		log.Fatalf("Cloudinary init failed: %v", err)
	}
	handler := handlers.NewProductHandler(service, cloudinaryService, imageService)

	routes := router.Group("/products")
	{
		routes.Use(auth.AuthMiddleware())
		routes.POST("/product", auth.RoleMiddleware("Admin"), handler.CreateProduct)
		routes.PUT("/product/:id", auth.RoleMiddleware("Admin"), handler.UpdateProduct)
		routes.DELETE("/product/:id", auth.RoleMiddleware("Admin"), handler.DeleteProduct)
		routes.POST("product/image", auth.RoleMiddleware("Admin"), handler.ImageUpload)
		routes.GET("/product/:id", handler.GetProductByID)
		routes.GET("", handler.GetAllProducts)
	}
}
