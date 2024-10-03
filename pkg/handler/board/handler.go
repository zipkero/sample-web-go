package board

// mongodb 사용

import (
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/internal/entity"
	"github.com/zipkero/sample-web-go/internal/service"
	"github.com/zipkero/sample-web-go/pkg/dto"
	"github.com/zipkero/sample-web-go/pkg/handler"
)

func FindAll(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		boards, err := boardService.FindAll(ctx)
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"data": boards,
		})
	}
}

func FindOne(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		board, err := boardService.FindOne(ctx, c.Param("id"))
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"data": board,
		})
	}
}

func Insert(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var board dto.BoardInsert

		if err := c.ShouldBind(&board); err != nil {
			handler.AbortWith(c, 400, "wrong request board data")
			return
		}

		id, err := boardService.InsertOne(ctx, board.ToEntity())
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(201, gin.H{
			"id": id,
		})
	}
}

func UpdateOne(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var board dto.BoardUpdate

		if err := c.ShouldBind(&board); err != nil {
			handler.AbortWith(c, 400, "wrong request board data")
			return
		}

		boardEntity, err := board.ToEntity()
		if err != nil {
			handler.AbortWith(c, 400, "modify entity error")
			return
		}

		err = boardService.UpdateOne(ctx, c.Param("id"), boardEntity)
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"data": boardEntity,
		})
	}
}

func DeleteOne(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		err := boardService.DeleteOne(ctx, c.Param("id"))
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}
		c.Status(200)
	}
}

func InsertMany(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var boards []dto.BoardInsert

		ctx := c.Request.Context()

		if err := c.ShouldBind(&boards); err != nil {
			handler.AbortWith(c, 400, "Bad Request")
			return
		}

		var entities []*entity.Board
		for _, board := range boards {
			entities = append(entities, board.ToEntity())
		}

		ids, err := boardService.InsertMany(ctx, entities)
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"data": ids,
		})
	}
}

func UpdateMany(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ctx := c.Request.Context()

		c.JSON(200, gin.H{
			"message": "Update Many",
		})
	}
}

func DeleteMany(boardService *service.BoardService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var boardIds dto.BoardDeleteMany
		if err := c.ShouldBind(&boardIds); err != nil {
			handler.AbortWith(c, 400, "Bad Request")
			return
		}

		var boardIdsString []string
		for _, boardId := range boardIds.IDs {
			boardIdsString = append(boardIdsString, boardId)
		}

		err := boardService.DeleteMany(ctx, boardIdsString)
		if err != nil {
			handler.AbortWith(c, 500, err.Error())
			return
		}

		c.JSON(200, gin.H{
			"data": boardIdsString,
		})
	}
}
