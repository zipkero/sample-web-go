package main

import (
	"flag"
	"fmt"
	"github.com/zipkero/sample-web-go/pkg/router"
	"github.com/zipkero/sample-web-go/pkg/server"
)

func main() {
	env := flag.String("env", "local", "환경설정")
	flag.Parse()

	svr, err := server.NewServer(fmt.Sprintf("config.%s.toml", *env))
	if err != nil {
		panic(err)
	}

	svr.UseRouter(router.RegisterRoute)

	if err := svr.Run(); err != nil {
		panic(err)
	}
}
