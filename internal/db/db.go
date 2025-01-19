package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/meyanksingh/vlink-backend/internal/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Error Loading the ENV File")
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		fmt.Println("Database connection established successfully!")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database models: %v", err)
	}
	return DB
}
