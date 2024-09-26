package board

// mongodb 사용

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Find All",
	})
}

func FindOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Find One: %s", c.Param("id")),
	})
}

func Insert(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Insert",
	})
}

func UpdateOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update",
	})
}

func DeleteOne(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}

func InsertMany(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Insert Many",
	})
}

func UpdateMany(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Many",
	})
}

func DeleteMany(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Many",
	})
}
