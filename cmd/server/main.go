package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	database "github.com/meyanksingh/vlink-backend/internal/db"
	routes "github.com/meyanksingh/vlink-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found - using environment variables from docker-compose")
	}

	database.ConnectDB()
	port := os.Getenv("PORT")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // specify frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.AuthRoutes(router)

	router.Run("0.0.0.0:" + port)

}
