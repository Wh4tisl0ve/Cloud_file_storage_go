package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string `env:"env"`
	HttpServer Server
	Database   Database
}

type Server struct {
	Host string `env:"HTTP_SERVER_HOST" envDefault:"localhost" required:"true"`
	Port int    `env:"HTTP_SERVER_PORT" envDefault:"8082" required:"true"`
}

type Database struct {
	Host     string `env:"DATABASE_HOST" required:"true"`
	Port     int    `env:"DATABASE_PORT" required:"true"`
	Username string `env:"DATABASE_USERNAME" required:"true"`
	Password string `env:"DATABASE_PASSWORD" required:"true"`
	Name     string `env:"DATABASE_NAME" required:"true"`
}

func MustLoad() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %s", err.Error())
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("Error parse config struct: %s", err.Error())
	}

	return &cfg, nil
}
