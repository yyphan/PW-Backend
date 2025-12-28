package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Admin-Secret")

		if token != os.Getenv("ADMIN_SECRET") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Who are you??!"})
			return
		}

		c.Next()
	}
}
