package main

import (
	"log"

	"github.com/dostonshernazarov/doctor-appointment/config"
	"github.com/dostonshernazarov/doctor-appointment/internal/app"
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
