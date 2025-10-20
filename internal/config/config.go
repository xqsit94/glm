package config

import (
	"encoding/json"
	"fmt"
	"os"

	"glm/pkg/paths"
)

type Config struct {
	AnthropicAuthToken string `json:"anthropic_auth_token"`
	DefaultModel       string `json:"default_model,omitempty"`
}

type ClaudeSettings struct {
	Env struct {
		AnthropicBaseURL   string `json:"ANTHROPIC_BASE_URL"`
		AnthropicAuthToken string `json:"ANTHROPIC_AUTH_TOKEN"`
		AnthropicModel     string `json:"ANTHROPIC_MODEL"`
	} `json:"env"`
}

func Load() (*Config, error) {
	configPath := paths.GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}

func Save(config *Config) error {
	configDir := paths.GetConfigDir()

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	configPath := paths.GetConfigPath()
	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

func LoadClaudeSettings() (*ClaudeSettings, error) {
	settingsPath := paths.GetClaudeSettingsPath()

	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("GLM is not enabled. Run 'glm enable' first")
	}

	data, err := os.ReadFile(settingsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read settings file: %v", err)
	}

	var settings ClaudeSettings
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse settings file: %v", err)
	}

	return &settings, nil
}

func SaveClaudeSettings(settings *ClaudeSettings) error {
	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %v", err)
	}

	settingsPath := paths.GetClaudeSettingsPath()
	if err := os.WriteFile(settingsPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write settings file: %v", err)
	}

	return nil
}
