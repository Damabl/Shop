package main

import (
	"Shop/internal/models"
	"Shop/internal/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(postgres.Open("postgres://myuser:mypassword@localhost:5444/mydatabase?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	err = db.AutoMigrate(&models.Product{}, &models.Cart{}, &models.CartItem{}, &models.User{})
	if err != nil {
		log.Fatal("Error on migrating to the DB", err)
	}
	r := gin.Default()
	routes.SetupUserRoutes(r, db)
	routes.SetupProductRoutes(r, db)
	routes.SetupCartRoutes(r, db)
	r.Run(":8080")
}
