package paths

import (
	"os"
	"path/filepath"
)

func GetClaudeDir() string {
	return filepath.Join(os.Getenv("HOME"), ".claude")
}

func GetClaudeSettingsPath() string {
	return filepath.Join(GetClaudeDir(), "settings.json")
}

func GetConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".glm")
}

func GetConfigPath() string {
	return filepath.Join(GetConfigDir(), "config.json")
}
