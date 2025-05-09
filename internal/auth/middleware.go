package auth

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}
		claims, err := ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
