package config

import (
	"log"
	"os"
	"path/filepath"
)

// Concats Hopper executable directory and file
func InExDir(file string) (string, error) {
	// get the executable path
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(filepath.Dir(ex), file), nil
}
