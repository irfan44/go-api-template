package config

import (
	"os"

	"github.com/irfan44/go-api-template/pkg/constants"
)

type Config struct {
	Http httpConfig
}

type httpConfig struct {
	Port string
	Host string
}

func NewConfig() Config {
	cfg := Config{
		Http: httpConfig{
			Port: os.Getenv(constants.HTTP_PORT),
			Host: os.Getenv(constants.HTTP_HOST),
		},
	}

	return cfg
}
