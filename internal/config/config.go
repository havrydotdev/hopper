package config

import (
	"log/slog"
	"os"
)

type Server struct {
	Port uint `toml:"port"`
}

type Logging struct {
	Level string `toml:"level" comment:"Logging level. Must be one of DEBUG, INFO or ERROR"`
}

type Motd struct {
	Description string `toml:"description"`
	FaviconPath string `toml:"favicon_path,commented" comment:"Path to the PNG favicon. Must be 64x64px."`
	MaxPlayers  uint   `toml:"max_players"`
}

type Config struct {
	Motd    Motd    `toml:"motd"`
	Logging Logging `toml:"logging"`
	Server  Server  `toml:"server"`
}

func Default() *Config {
	cfg := new(Config)

	// server options
	cfg.Server.Port = 25565

	// logger options
	cfg.Logging.Level = "INFO"

	// MOTD options
	cfg.Motd.MaxPlayers = 20
	cfg.Motd.Description = "Hopper Minecraft Server"
	cfg.Motd.FaviconPath = "favicon.png"

	return cfg
}

func (c *Config) InitLogger() {
	handler := slog.NewTextHandler(os.Stdout, c.LoggerOptions())
	slog.SetDefault(slog.New(handler))
}

func (c *Config) LoggerOptions() *slog.HandlerOptions {
	opts := new(slog.HandlerOptions)
	opts.Level = c.LogLevel()

	return opts
}

func (c *Config) LogLevel() slog.Leveler {
	switch c.Logging.Level {
	case "DEBUG":
		return slog.LevelDebug
	case "WARN":
		return slog.LevelInfo
	case "ERROR":
		return slog.LevelError
	}

	return slog.LevelInfo
}
