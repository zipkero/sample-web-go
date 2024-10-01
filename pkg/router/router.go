package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/internal/service"
	"github.com/zipkero/sample-web-go/pkg/handler/board"
	"github.com/zipkero/sample-web-go/pkg/handler/user"
	"github.com/zipkero/sample-web-go/pkg/middleware"
	"github.com/zipkero/sample-web-go/pkg/server"
)

func RegisterRoute(svr *server.Server) {
	svr.RegisterRoute("GET", "/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	svr.RegisterRoute("GET", "/board", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.FindAll(boardService)(c)
	})
	svr.RegisterRoute("GET", "/board/:id", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.FindOne(boardService)(c)
	})

	svr.RegisterRoute("POST", "/user/login", user.Login)
	svr.RegisterRoute("POST", "/user/register", user.Register)

	authGroup := svr.RegisterGroup("/", middleware.AuthMiddleware())
	authGroup.POST("/board", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.Insert(boardService)(c)
	})
	authGroup.PATCH("/board/:id", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.UpdateOne(boardService)(c)
	})
	authGroup.DELETE("/board/:id", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.DeleteOne(boardService)(c)
	})
	authGroup.POST("/boards", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.InsertMany(boardService)(c)
	})
	authGroup.PATCH("/boards", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.UpdateMany(boardService)(c)
	})
	authGroup.DELETE("/boards", func(c *gin.Context) {
		boardService := service.NewBoardService(svr.MongoProvider)

		board.DeleteMany(boardService)(c)
	})

	authGroup.POST("/user/logout", user.Logout)
}
