package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/drshashwat/coolstream/server/CoolStreamMovieServer/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.GetAccessToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
