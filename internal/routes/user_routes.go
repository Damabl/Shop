package routes

import (
	"Shop/internal/auth"
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
	authHandler := auth.NewAuthHandler(userService)
	userRoutes := r.Group("/users")
	{
		userRoutes.Use(auth.AuthMiddleware())
		userRoutes.GET("/:id", userHandler.GetUserProfile)
		userRoutes.GET("/", auth.RoleMiddleware("Admin"), userHandler.GetAllUsers)
	}
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}
}
