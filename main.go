package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var rootCmd = &cobra.Command{
	Use:   "glm",
	Short: "GLM Claude settings management CLI",
	Long:  "A CLI tool to enable/disable GLM settings for Claude",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDefaultAction()
	},
}

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable GLM settings for Claude",
	Long:  "Configure Claude to use GLM model via BigModel API",
	RunE: func(cmd *cobra.Command, args []string) error {
		model, _ := cmd.Flags().GetString("model")
		return enableGLM(model)
	},
}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable GLM settings for Claude",
	Long:  "Remove GLM configuration and restore default Claude settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		return disableGLM()
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change GLM model settings",
	Long:  "Update the ANTHROPIC_MODEL in existing GLM configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		model, _ := cmd.Flags().GetString("model")
		if model == "" {
			return fmt.Errorf("model flag is required")
		}
		return setGLMModel(model)
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install tools",
	Long:  "Install various tools and dependencies",
}

var installClaudeCmd = &cobra.Command{
	Use:   "claude",
	Short: "Install Claude Code",
	Long:  "Install Claude Code using npm",
	RunE: func(cmd *cobra.Command, args []string) error {
		return installClaude()
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage authentication token",
	Long:  "Manage your Anthropic authentication token",
}

var tokenSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set authentication token",
	Long:  "Set your Anthropic authentication token interactively",
	RunE: func(cmd *cobra.Command, args []string) error {
		return setAuthToken()
	},
}

var tokenShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current token",
	Long:  "Display the current authentication token (masked)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return showAuthToken()
	},
}

var tokenClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear authentication token",
	Long:  "Remove the stored authentication token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return clearAuthToken()
	},
}

func init() {
	enableCmd.Flags().StringP("model", "m", "glm-4.5", "GLM model to use (default: glm-4.5)")
	setCmd.Flags().StringP("model", "m", "", "GLM model to set (required)")
	setCmd.MarkFlagRequired("model")

	installCmd.AddCommand(installClaudeCmd)
	tokenCmd.AddCommand(tokenSetCmd)
	tokenCmd.AddCommand(tokenShowCmd)
	tokenCmd.AddCommand(tokenClearCmd)

	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(tokenCmd)
}

func enableGLM(model string) error {
	claudeDir := filepath.Join(os.Getenv("HOME"), ".claude")
	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	token, err := getAuthToken()
	if err != nil {
		return err
	}

	settingsPath := filepath.Join(claudeDir, "settings.json")
	content := fmt.Sprintf(`{
  "env": {
      "ANTHROPIC_BASE_URL": "https://open.bigmodel.cn/api/anthropic",
      "ANTHROPIC_AUTH_TOKEN": "%s",
      "ANTHROPIC_MODEL": "%s"
  }
}`, token, model)

	if err := os.WriteFile(settingsPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write settings file: %v", err)
	}

	fmt.Printf("Claude settings have been configured successfully with model: %s\n", model)
	return nil
}

func disableGLM() error {
	claudeDir := filepath.Join(os.Getenv("HOME"), ".claude")
	settingsPath := filepath.Join(claudeDir, "settings.json")

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


type ClaudeSettings struct {
	Env struct {
		AnthropicBaseURL   string `json:"ANTHROPIC_BASE_URL"`
		AnthropicAuthToken string `json:"ANTHROPIC_AUTH_TOKEN"`
		AnthropicModel     string `json:"ANTHROPIC_MODEL"`
	} `json:"env"`
}

type Config struct {
	AnthropicAuthToken string `json:"anthropic_auth_token"`
	DefaultModel       string `json:"default_model,omitempty"`
}

func setGLMModel(model string) error {
	claudeDir := filepath.Join(os.Getenv("HOME"), ".claude")
	settingsPath := filepath.Join(claudeDir, "settings.json")

	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		return fmt.Errorf("GLM is not enabled. Run 'glm enable' first")
	}

	data, err := os.ReadFile(settingsPath)
	if err != nil {
		return fmt.Errorf("failed to read settings file: %v", err)
	}

	var settings ClaudeSettings
	if err := json.Unmarshal(data, &settings); err != nil {
		return fmt.Errorf("failed to parse settings file: %v", err)
	}

	settings.Env.AnthropicModel = model


	updatedData, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %v", err)
	}

	if err := os.WriteFile(settingsPath, updatedData, 0644); err != nil {
		return fmt.Errorf("failed to write settings file: %v", err)
	}

	fmt.Printf("GLM model has been updated to: %s\n", model)
	return nil
}

func installClaude() error {
	if !isNpmAvailable() {
		fmt.Println("❌ npm is not available on your system.")
		fmt.Println("📦 To install Claude Code, you need Node.js and npm.")
		fmt.Println("🔗 Please install Node.js from: https://nodejs.org/")
		fmt.Println("💡 After installing Node.js, npm will be available automatically.")
		fmt.Println("🔄 Then run 'glm install claude' again.")
		return fmt.Errorf("npm not found")
	}

	fmt.Println("📦 Installing Claude Code...")
	fmt.Println("🔄 Running: npm install -g @anthropic-ai/claude-code")

	cmd := exec.Command("npm", "install", "-g", "@anthropic-ai/claude-code")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install Claude Code: %v", err)
	}

	fmt.Println("✅ Claude Code has been installed successfully!")
	fmt.Println("🚀 You can now use 'claude' command from anywhere.")
	return nil
}

func isNpmAvailable() bool {
	_, err := exec.LookPath("npm")
	return err == nil
}

func getConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), ".glm")
}

func getConfigPath() string {
	return filepath.Join(getConfigDir(), "config.json")
}

func loadConfig() (*Config, error) {
	configPath := getConfigPath()
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

func saveConfig(config *Config) error {
	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %v", err)
	}

	configPath := getConfigPath()
	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %v", err)
	}

	return nil
}

func setAuthToken() error {
	fmt.Print("Enter your Anthropic API token: ")

	token, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("failed to read token: %v", err)
	}
	fmt.Println()

	tokenStr := strings.TrimSpace(string(token))
	if tokenStr == "" {
		return fmt.Errorf("token cannot be empty")
	}

	config, err := loadConfig()
	if err != nil {
		return err
	}

	config.AnthropicAuthToken = tokenStr
	if config.DefaultModel == "" {
		config.DefaultModel = "glm-4.5"
	}

	if err := saveConfig(config); err != nil {
		return err
	}

	fmt.Println("✅ Authentication token has been saved successfully!")
	return nil
}

func showAuthToken() error {
	token, err := getAuthToken()
	if err != nil {
		return err
	}

	if len(token) > 8 {
		masked := token[:4] + strings.Repeat("*", len(token)-8) + token[len(token)-4:]
		fmt.Printf("Current token: %s\n", masked)
	} else {
		fmt.Println("Current token: ****")
	}
	return nil
}

func clearAuthToken() error {
	configPath := getConfigPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("No token found to clear.")
		return nil
	}

	if err := os.Remove(configPath); err != nil {
		return fmt.Errorf("failed to remove config file: %v", err)
	}

	configDir := getConfigDir()
	if entries, err := os.ReadDir(configDir); err == nil && len(entries) == 0 {
		os.Remove(configDir)
	}

	fmt.Println("✅ Authentication token has been cleared successfully!")
	return nil
}

func getAuthToken() (string, error) {
	if token := os.Getenv("ANTHROPIC_AUTH_TOKEN"); token != "" {
		return token, nil
	}

	config, err := loadConfig()
	if err != nil {
		return "", err
	}
	if config.AnthropicAuthToken != "" {
		return config.AnthropicAuthToken, nil
	}


	fmt.Println("🔐 No authentication token found.")
	fmt.Print("Would you like to set up your token now? (y/n): ")

	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) == "y" || strings.ToLower(response) == "yes" {
		if err := setAuthToken(); err != nil {
			return "", err
		}
		return getAuthToken()
	}

	return "", fmt.Errorf("authentication token is required. Use 'glm token set' to configure it")
}


func runDefaultAction() error {
	fmt.Println("🚀 Running default GLM action...")

	fmt.Println("📝 Enabling GLM...")
	if err := enableGLM("glm-4.5"); err != nil {
		return fmt.Errorf("failed to enable GLM: %v", err)
	}

	fmt.Println("🎯 Starting Claude Code...")

	if _, err := exec.LookPath("claude"); err != nil {
		fmt.Println("❌ Claude Code is not installed.")
		fmt.Println("💡 Run 'glm install claude' first to install Claude Code.")
		return fmt.Errorf("claude command not found")
	}


	cmd := exec.Command("claude")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run claude: %v", err)
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}