package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/internal/config"
	"github.com/zipkero/sample-web-go/pkg/middleware"
)

type Server struct {
	config *config.Config
	engine *gin.Engine
}

func NewServer(path string) (*Server, error) {
	svr := new(Server)
	cfg, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	svr.config = cfg

	return svr, nil
}

func (s *Server) Register(method, path string, handler gin.HandlerFunc) {
	s.engine.Handle(method, path, middleware.AuthMiddleware(), handler)
}

func (s *Server) Run() error {
	return s.engine.Run(s.config.App.Addr)
}
