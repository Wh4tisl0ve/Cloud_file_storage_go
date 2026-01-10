package main

import (
	"log"

	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/app"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/config"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	
	app.Run(cfg)
}
