package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/zipkero/sample-web-go/internal/config"
)

type RedisProvider struct {
	client *redis.Client
}

func NewRedisProvider(config *config.Config) (*RedisProvider, error) {
	opts, err := redis.ParseURL(config.Redis.URI)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	return &RedisProvider{client}, nil
}
