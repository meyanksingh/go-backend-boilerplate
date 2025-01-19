package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meyanksingh/vlink-backend/internal/app/controllers"
	middleware "github.com/meyanksingh/vlink-backend/internal/middleware"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	Auth := incomingRoutes.Group("/auth")
	{
		Auth.POST("/register", controller.Register)
		Auth.POST("/login", controller.Login)
		Auth.GET("/", middleware.JWTAuthMiddleware(), controller.Home)
	}

}
