package board

// mongodb 사용

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/internal/service"
	"github.com/zipkero/sample-web-go/pkg/dto"
	"github.com/zipkero/sample-web-go/pkg/handler"
	"strconv"
)

func FindAll(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Find All",
		})
	}
}

func FindOne(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			handler.AbortWith(c, 400, "id must be integer")
			return
		}

		board, err := boardService.FindOne(ctx, id)
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"data": board,
		})
	}
}

func Insert(c *gin.Context) {
	var board dto.BoardInsert

	if err := c.ShouldBind(&board); err != nil {
		handler.AbortWith(c, 400, "Bad Request")
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Insert: %s", board.Title),
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
	var boards []dto.BoardInsert

	if err := c.ShouldBind(&boards); err != nil {
		handler.AbortWith(c, 400, "Bad Request")
		return
	}

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
