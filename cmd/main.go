package main

import (
	"log"
	"log/slog"

	"github.com/gavrylenkoIvan/hopper/internal/config"
	"github.com/gavrylenkoIvan/hopper/internal/server"
)

func main() {
	// Get executable's directory
	configPath, err := config.InExDir("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Read(configPath)
	if err != nil {
		log.Fatal("Failed to read config: ", err)
	}

	// Initialize hopper logger
	cfg.InitLogger()

	slog.Info("Config is valid")

	hopper, err := server.New(cfg, nil)
	if err != nil {
		log.Fatal("Failed to create server instance: ", err)
	}

	err = hopper.Listen()
	if err != nil {
		log.Fatal("Failed to start Hopper server: ", err)
	}
}
