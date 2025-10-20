package glm

import (
	"fmt"
	"os"

	"github.com/xqsit94/glm/internal/config"
	"github.com/xqsit94/glm/pkg/paths"
)

func Enable(model, token string) error {
	claudeDir := paths.GetClaudeDir()

	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	settings := &config.ClaudeSettings{}
	settings.Env.AnthropicBaseURL = "https://open.bigmodel.cn/api/anthropic"
	settings.Env.AnthropicAuthToken = token
	settings.Env.AnthropicModel = model

	if err := config.SaveClaudeSettings(settings); err != nil {
		return err
	}

	fmt.Printf("Claude settings have been configured successfully with model: %s\n", model)
	return nil
}

func Disable() error {
	claudeDir := paths.GetClaudeDir()
	settingsPath := paths.GetClaudeSettingsPath()

	if _, err := os.Stat(settingsPath); err == nil {
		if err := os.Remove(settingsPath); err != nil {
			return fmt.Errorf("failed to remove settings file: %v", err)
		}
		fmt.Println("Claude settings file has been removed.")
	} else {
		fmt.Println("Claude settings file not found.")
	}

	if entries, err := os.ReadDir(claudeDir); err == nil {
		if len(entries) == 0 {
			if err := os.Remove(claudeDir); err != nil {
				return fmt.Errorf("failed to remove directory: %v", err)
			}
			fmt.Println("Empty .claude directory has been removed.")
		} else {
			fmt.Println(".claude directory contains other files and was not removed.")
		}
	}

	fmt.Println("Cleanup completed.")
	return nil
}

func SetModel(model string) error {
	settings, err := config.LoadClaudeSettings()
	if err != nil {
		return err
	}

	settings.Env.AnthropicModel = model

	if err := config.SaveClaudeSettings(settings); err != nil {
		return err
	}

	fmt.Printf("GLM model has been updated to: %s\n", model)
	return nil
}
