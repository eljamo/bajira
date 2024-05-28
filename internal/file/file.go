package file

import (
	"path/filepath"

	"github.com/eljamo/bajira/internal/config"
)

func GetBajiraConfigFile() string {
	dir, err := GetConfigDirectory()
	if err != nil {
		return ""
	}

	return filepath.Join(dir, config.BajiraFileNameConfig)
}
