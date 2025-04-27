package db

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var DB *gorm.DB

func InitDB() {
	dbName := os.Getenv("DATABASE_NAME")
	sqlDB, err := sql.Open(dbName, os.Getenv("APP_POSTGRES_URL"))
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Error creating migrate driver: ", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working dir: %v", err)
	}
	migrationsPath := filepath.Join(cwd, "internal", "db", "migrations")
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath)
	m, err := migrate.NewWithDatabaseInstance(migrationsURL, dbName, driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}
	err = m.Up()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening GORM connection: ", err)
	}

	DB = gormDB
}
