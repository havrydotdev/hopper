package main

import (
	"log"
	"log/slog"

	"havry.dev/havry/hopper/internal/config"
	"havry.dev/havry/hopper/internal/server"
)

func main() {
	// get executable's directory
	configPath, err := config.InExDir("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Read(configPath)
	if err != nil {
		log.Fatal("Failed to read config: ", err)
	}

	// initialize hopper logger
	cfg.InitLogger()

	slog.Info("Config is valid")

	err = server.New(cfg, nil).Listen()
	if err != nil {
		log.Fatal("Failed to start Hopper server: ", err)
	}
}
