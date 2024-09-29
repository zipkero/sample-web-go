package handler

import "github.com/gin-gonic/gin"

var defaultFailMessages = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	500: "Internal Server Error",
}

func AbortWith(c *gin.Context, code int, message string) {
	if message == "" {
		if msg, ok := defaultFailMessages[code]; ok {
			message = msg
		} else {
			message = "unknown code error"
		}
	}

	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}
