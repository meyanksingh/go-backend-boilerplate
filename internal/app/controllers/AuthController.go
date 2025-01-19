package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meyanksingh/vlink-backend/internal/app/models"
	"github.com/meyanksingh/vlink-backend/internal/app/repository"
	database "github.com/meyanksingh/vlink-backend/internal/db"
	"github.com/meyanksingh/vlink-backend/pkg/utils"
)

func Register(c *gin.Context) {
	var requestBody struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "yes dadyInvalid request data"})
		return
	}

	emailExists, err := repository.CheckEmailExists(requestBody.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking email"})
		return
	}
	if emailExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already in use"})
		return
	}

	user, err := repository.CreateUser(requestBody.FirstName, requestBody.LastName, requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
		"email":   user.Email,
	})
}

func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	user, err := repository.AuthenticateUser(requestBody.Email, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}

func Home(c *gin.Context) {
	userIDStr, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid User ID"})
		return
	}

	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the protected home page!",
		"email":   user.Email,
		"name":    user.FirstName + " " + user.LastName,
	})
}
