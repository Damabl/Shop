package main

import (
	"Shop/internal/cloud"
	"Shop/internal/db"
	"Shop/internal/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	db.InitDB()
	r := gin.Default()
	cloud.NewCloudinaryService()
	routes.SetupCartRoutes(r, db.DB)
	routes.SetupProductRoutes(r, db.DB)
	routes.SetupUserRoutes(r, db.DB)
	srv := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: r,
	}
	srv.ListenAndServe()
}
