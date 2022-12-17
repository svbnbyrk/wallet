package main

import (
	"log"

	"github.com/svbnbyrk/wallet/config"
	"github.com/svbnbyrk/wallet/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
