package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/meyanksingh/vlink-backend/pkg/utils"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from token if it's present
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = tokenString[len("Bearer "):]
		}

		// Validate the token
		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims and add them to the context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Attach user_id to the context for use in protected routes
		userID, ok := claims["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token user ID"})
			c.Abort()
			return
		}
		c.Set("user_id", userID) // Set the user ID in context
		c.Next()                 // Proceed to the next handler
	}
}
