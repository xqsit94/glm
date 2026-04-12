package paths

import (
	"os"
	"path/filepath"
)

func GetConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".glm")
}

func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.json")
}
