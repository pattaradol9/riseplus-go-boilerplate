package main

import (
	"log"
	"riseplus-go-boilerplate/internal/config"
	"riseplus-go-boilerplate/internal/server"

	_ "go.uber.org/automaxprocs"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app := server.New(cfg)
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
