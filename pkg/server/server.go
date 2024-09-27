package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/internal/config"
	"github.com/zipkero/sample-web-go/internal/db"
	"github.com/zipkero/sample-web-go/pkg/middleware"
)

type Server struct {
	config *config.Config
	engine *gin.Engine

	mongoProvider *db.MongoProvider
	redisProvider *db.RedisProvider
}

func NewServer(path string) (*Server, error) {
	svr := new(Server)

	svr.engine = gin.Default()

	cfg, err := config.NewConfig(path)
	if err != nil {
		return nil, err
	}

	svr.config = cfg
	svr.mongoProvider, err = db.NewMongoProvider(cfg)
	if err != nil {
		return nil, err
	}

	svr.redisProvider, err = db.NewRedisProvider(cfg)
	if err != nil {
		return nil, err
	}

	return svr, nil
}

func (s *Server) Register(method, path string, handlers ...gin.HandlerFunc) {
	handlers = append(handlers, middleware.AuthMiddleware())

	s.engine.Handle(method, path, handlers...)
}

func (s *Server) Run() error {
	return s.engine.Run(s.config.App.Addr)
}
