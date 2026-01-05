package app

import (
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/config"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/logger"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/postgres"
)

func Run(cfg *config.Config) {
	// logger setup
	envConfig := cfg.Env
	logger := logger.SetupLogger(envConfig)

	// DB setup
	dbConfig := cfg.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	postgres, err := postgres.New(dsn)
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("✅ Connected to PostgreSQL successfully!")
	}
	defer postgres.Close()
}
