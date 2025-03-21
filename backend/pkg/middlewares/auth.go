package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Set("user_type", claims["UserType"])
		c.Set("email", claims["email"])
		c.Next()
	}
}