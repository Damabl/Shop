package routes

import (
	"Shop/internal/auth"
	"Shop/internal/handlers"
	"Shop/internal/repositories"
	"Shop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCartRoutes(r *gin.Engine, db *gorm.DB) {
	cartRepo := repositories.NewCartRepository(db)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService)

	cartRoutes := r.Group("/cart")
	{
		cartRoutes.Use(auth.AuthMiddleware())
		cartRoutes.GET("/user/:user_id", cartHandler.GetCart)
		cartRoutes.POST("/user/:user_id/product/:product_id", cartHandler.AddToCart)
		cartRoutes.PUT("/item/:cart_item_id", cartHandler.UpdateCartItem)
		cartRoutes.DELETE("/item/:cart_item_id", cartHandler.RemoveFromCart)
		cartRoutes.DELETE("/user/:user_id/clear", cartHandler.ClearCart)
	}
}
