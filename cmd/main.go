package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/pkg/handler/board"
	"github.com/zipkero/sample-web-go/pkg/handler/user"
	"github.com/zipkero/sample-web-go/pkg/server"
)

func main() {
	env := flag.String("env", "local", "환경설정")
	flag.Parse()

	svr, err := server.NewServer(fmt.Sprintf("config.%s.toml", *env))
	if err != nil {
		panic(err)
	}

	svr.Register("GET", "/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	svr.Register("GET", "/board", board.FindAll)
	svr.Register("POST", "/board", board.Insert)
	svr.Register("GET", "/board/:id", board.FindOne)
	svr.Register("PATCH", "/board/:id", board.UpdateOne)
	svr.Register("DELETE", "/board/:id", board.DeleteOne)

	svr.Register("POST", "/boards", board.InsertMany)
	svr.Register("PATCH", "/boards", board.UpdateMany)
	svr.Register("DELETE", "/boards", board.DeleteMany)

	svr.Register("POST", "/user/login", user.Login)
	svr.Register("POST", "/user/logout", user.Logout)
	svr.Register("POST", "/user/register", user.Register)

	if err := svr.Run(); err != nil {
		panic(err)
	}
}
