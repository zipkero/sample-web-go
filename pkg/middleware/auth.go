package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Auth Middleware")

		c.Next()
	}
}
