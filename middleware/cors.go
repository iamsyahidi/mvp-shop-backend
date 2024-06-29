package middleware

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		wHead := c.Writer.Header()
		wHead.Set("Access-Control-Allow-Origin", "*")
		wHead.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
		wHead.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin, Cookie, Signature, Timestamp")
		wHead.Set("Access-Control-Allow-Credentials", "true")
		wHead.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		wHead.Set("Cache-Control", "no-store")
		wHead.Set("X-Content-Type-Options", "nosniff")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
