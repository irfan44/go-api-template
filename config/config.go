package config

import (
	"os"

	"github.com/irfan44/go-example/pkg/constants"
)

type (
	Config struct {
		Http     httpConfig
		Postgres PostgresConfig
	}

	httpConfig struct {
		Port string
		Host string
	}

	PostgresConfig struct {
		Port     string
		Host     string
		User     string
		Password string
		DBName   string
	}
)

func NewConfig() Config {
	cfg := Config{
		Http: httpConfig{
			Port: os.Getenv(constants.HTTP_PORT),
			Host: os.Getenv(constants.HTTP_HOST),
		},
		Postgres: PostgresConfig{
			Port:     os.Getenv(constants.DB_PORT),
			Host:     os.Getenv(constants.DB_HOST),
			User:     os.Getenv(constants.DB_USER),
			Password: os.Getenv(constants.DB_PASSWORD),
			DBName:   os.Getenv(constants.DB_NAME),
		},
	}

	return cfg
}
