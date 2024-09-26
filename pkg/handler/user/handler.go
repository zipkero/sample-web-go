package user

// redis 사용

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login",
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Logout",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register",
	})
}
