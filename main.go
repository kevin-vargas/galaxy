package main

import (
	"log"

	"github.com/kevin-vargas/galaxy/internal/app"

	"github.com/kevin-vargas/galaxy/config"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
