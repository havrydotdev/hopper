package config

import (
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func CreateAndWrite(cfg *Config, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return WriteRaw(cfg, file)
}

func Write(cfg *Config, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return WriteRaw(cfg, file)
}

func WriteRaw(cfg *Config, w io.Writer) error {
	encoder := toml.NewEncoder(w)
	err := encoder.SetIndentTables(true).Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
