package config

import (
	"log"

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

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Error parse config struct")
	}

	return &cfg
}
