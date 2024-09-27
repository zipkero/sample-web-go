package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zipkero/sample-web-go/internal/config"
	"github.com/zipkero/sample-web-go/internal/db"
)

type Server struct {
	config *config.Config
	engine *gin.Engine

	mongoProvider *db.MongoProvider
	redisProvider *db.RedisProvider
}

func NewServer(path string) (*Server, error) {
	svr := new(Server)

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

	svr.engine = gin.Default()

	return svr, nil
}

func (s *Server) UseRouter(router func(s *Server)) {
	router(s)
}

func (s *Server) RegisterGroup(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return s.engine.Group(path, handlers...)
}

func (s *Server) RegisterRoute(method, path string, handlers ...gin.HandlerFunc) {
	s.engine.Handle(method, path, handlers...)
}

func (s *Server) Run() error {
	return s.engine.Run(s.config.App.Addr)
}
