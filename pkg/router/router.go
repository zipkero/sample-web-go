package router

import (
	"github.com/gin-gonic/gin"
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

	svr.RegisterRoute("GET", "/board", board.FindAll)
	svr.RegisterRoute("GET", "/board/:id", board.FindOne)

	svr.RegisterRoute("POST", "/user/login", user.Login)
	svr.RegisterRoute("POST", "/user/register", user.Register)

	authGroup := svr.RegisterGroup("/", middleware.AuthMiddleware())
	authGroup.POST("/board", board.Insert)
	authGroup.PATCH("/board/:id", board.UpdateOne)
	authGroup.DELETE("/board/:id", board.DeleteOne)
	authGroup.POST("/user/logout", user.Logout)
	authGroup.POST("/boards", board.InsertMany)
	authGroup.PATCH("/boards", board.UpdateMany)
	authGroup.DELETE("/boards", board.DeleteMany)
}
