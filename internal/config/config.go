package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	App struct {
		Addr string
	}
	Mongo struct {
		URI string
	}
	Redis struct {
		URI string
	}
}

func NewConfig(path string) (*Config, error) {
	c := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	if err := toml.NewDecoder(file).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}
