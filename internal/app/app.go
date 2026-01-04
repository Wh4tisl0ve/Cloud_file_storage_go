package app

import (
	"fmt"
	"log"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/config"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/pkg/postgres"
)

func Run(cfg *config.Config) {
	// DB config
	dbConfig := cfg.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	postgres, err := postgres.New(dsn)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Print("✅ Connected to PostgreSQL successfully!")
	}
	defer postgres.Close()
}
