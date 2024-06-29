package middleware

import (
	"mvp-shop-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a sample middleware for authentication and authorization using JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		decodes, err := JwtClaim(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.Response{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			})
			c.Abort()
			return
		}

		c.Set("customer", decodes)

		c.Next()
	}
}
