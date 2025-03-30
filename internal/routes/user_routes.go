package routes

import (
	"Shop/internal/handlers"
	"Shop/internal/repositories"
	"Shop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.GET("/:id", userHandler.GetUserProfile)
		userRoutes.GET("/", userHandler.GetAllUsers)
	}
}
