package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"glm/internal/glm"
	"glm/internal/token"

	"github.com/spf13/cobra"
)

const (
	version      = "1.1.0"
	defaultModel = "glm-4.6"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "glm",
		Short:   "GLM Claude settings management CLI",
		Long:    "A CLI tool to enable/disable GLM settings for Claude",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultAction()
		},
	}
}

func runDefaultAction() error {
	fmt.Println("🚀 Running default GLM action...")

	authToken, err := token.Get()
	if err != nil {
		return fmt.Errorf("failed to get authentication token: %v", err)
	}

	fmt.Println("📝 Enabling GLM...")
	if err := glm.Enable(defaultModel, authToken); err != nil {
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
