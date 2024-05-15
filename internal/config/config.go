package config

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
	cfg.Server.Port = 25565

	cfg.Logging.Level = "INFO"

	cfg.Motd.MaxPlayers = 20
	cfg.Motd.Description = "Hopper Minecraft Server"
	cfg.Motd.FaviconPath = "favicon.png"

	return cfg
}
