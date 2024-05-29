package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/toml"
)

func GetBajiraConfigFilePath() (string, error) {
	dir, err := GetConfigDirectory()
	if err != nil {
		return "", err
	}

	path := filepath.Join(dir, config.BajiraFileNameConfig)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return "", fmt.Errorf("failed to create file: %w", err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf(`failed to close file "%v": %v`, file, err)
			}
		}()
	}

	return path, nil
}

// GetBajiraConfig returns the configuration for Bajira.
func GetBajiraConfig() (*config.BajiraConfig, error) {
	path, err := GetBajiraConfigFilePath()
	if err != nil {
		return nil, err
	}

	// file is a toml file so decode it using toml.DecodeFromFile
	var config config.BajiraConfig
	err = toml.DecodeFromFile(path, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
