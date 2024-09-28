package handler

import "github.com/gin-gonic/gin"

var defaultFailMessages = map[int]string{
	400: "Bad Request",
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
