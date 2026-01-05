package main

import (
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/config"
	"github.com/Wh4tisl0ve/Cloud_file_storage_go/internal/app"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
