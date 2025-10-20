package token

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/xqsit94/glm/internal/config"
	"github.com/xqsit94/glm/pkg/paths"

	"golang.org/x/term"
)

const defaultModel = "glm-4.6"

func Get() (string, error) {
	if token := os.Getenv("ANTHROPIC_AUTH_TOKEN"); token != "" {
		return token, nil
	}

	cfg, err := config.Load()
	if err != nil {
		return "", err
	}
	if cfg.AnthropicAuthToken != "" {
		return cfg.AnthropicAuthToken, nil
	}

	fmt.Println("🔐 No authentication token found.")
	fmt.Print("Would you like to set up your token now? (y/n): ")

	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
		if err := Set(); err != nil {
			return "", err
		}
		return Get()
	}

	return "", fmt.Errorf("authentication token is required. Use 'glm token set' to configure it")
}

func Set() error {
	fmt.Print("Enter your Anthropic API token: ")

	tokenBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read token: %v", err)
	}
	fmt.Println()

	tokenStr := strings.TrimSpace(string(tokenBytes))
	if tokenStr == "" {
		return fmt.Errorf("token cannot be empty")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	cfg.AnthropicAuthToken = tokenStr
	if cfg.DefaultModel == "" {
		cfg.DefaultModel = defaultModel
	}

	if err := config.Save(cfg); err != nil {
		return err
	}

	fmt.Println("✅ Authentication token has been saved successfully!")
	return nil
}

func Show() error {
	token, err := Get()
	if err != nil {
		return err
	}

	// Mask token: show first 4 and last 4 chars
	if len(token) > 8 {
		masked := token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
		fmt.Printf("Current token: %s\n", masked)
	} else {
		fmt.Println("Current token: ****")
	}

	return nil
}

func Clear() error {
	configPath := paths.GetConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("No token found to clear.")
		return nil
	}

	if err := os.Remove(configPath); err != nil {
		return fmt.Errorf("failed to remove config file: %v", err)
	}

	configDir := paths.GetConfigDir()
	if entries, err := os.ReadDir(configDir); err == nil && len(entries) == 0 {
		os.Remove(configDir)
	}

	fmt.Println("✅ Authentication token has been cleared successfully!")
	return nil
}
