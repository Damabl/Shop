package routes

import (
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
		cartRoutes.GET("/:user_id", cartHandler.GetCart)
		cartRoutes.POST("/:user_id/product/:product_id", cartHandler.AddToCart)
		cartRoutes.PUT("/:cart_item_id", cartHandler.UpdateCartItem)
		cartRoutes.DELETE("/:cart_item_id", cartHandler.RemoveFromCart)
		cartRoutes.DELETE("/:user_id/clear", cartHandler.ClearCart)
	}
}
