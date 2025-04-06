package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// Open PostgreSQL connection
	sqlDB, err := sql.Open("postgres", os.Getenv("APP_POSTGRES_URL"))
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	// Create migration driver
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Error creating migrate driver: ", err)
	}

	// Initialize migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}

	// Check if migrations need to be applied
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		if err == migrate.ErrNilVersion {
			// This means no migrations have been applied yet
			fmt.Println("No migrations applied yet, applying now...")
		} else {
			log.Fatal("Error applying migrations: ", err)
		}
	}

	// Get current version (will only work after at least one migration)
	version, dirty, err := m.Version()
	if err == nil {
		fmt.Printf("Current migration version: %d (dirty: %v)\n", version, dirty)
	} else if err == migrate.ErrNilVersion {
		fmt.Println("No migrations have been applied yet")
	} else {
		log.Fatal("Error getting migration version: ", err)
	}

	// Initialize GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening GORM connection: ", err)
	}

	DB = gormDB
}
