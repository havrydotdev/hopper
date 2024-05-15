package config

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Read(path string) (*Config, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		cfg := Default()

		return cfg, CreateAndWrite(cfg, path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := new(Config)
	err = toml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
