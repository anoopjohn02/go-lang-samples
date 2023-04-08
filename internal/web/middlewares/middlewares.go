package middlewares

import (
	"com/anoop/examples/internal/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(tokenValidator *token.TokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenValidator.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
