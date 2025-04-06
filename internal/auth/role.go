package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have the required role"})
			c.Abort()
			return
		}
		c.Next()
	}
}
