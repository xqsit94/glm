package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/xqsit94/glm/internal/token"

	"github.com/spf13/cobra"
)

const (
	version      = "1.1.0"
	defaultModel = "glm-4.6"
)

func RootCmd() *cobra.Command {
	var model string

	cmd := &cobra.Command{
		Use:     "glm",
		Short:   "GLM Claude settings management CLI",
		Long:    "A CLI tool to launch Claude with GLM settings using temporary session-based configuration",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultAction(model)
		},
	}

	cmd.Flags().StringVarP(&model, "model", "m", defaultModel, "GLM model to use for this session")

	return cmd
}

func runDefaultAction(model string) error {
	fmt.Println("🚀 Launching Claude with GLM...")

	authToken, err := token.Get()
	if err != nil {
		return fmt.Errorf("failed to get authentication token: %v", err)
	}

	if _, err := exec.LookPath("claude"); err != nil {
		fmt.Println("❌ Claude Code is not installed.")
		fmt.Println("💡 Run 'glm install claude' first to install Claude Code.")
		return fmt.Errorf("claude command not found")
	}

	fmt.Printf("📝 Using model: %s\n", model)
	fmt.Println("🎯 Starting Claude Code with temporary GLM configuration...")

	cmd := exec.Command("claude")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		"ANTHROPIC_BASE_URL=https://open.bigmodel.cn/api/anthropic",
		"ANTHROPIC_AUTH_TOKEN="+authToken,
		"ANTHROPIC_MODEL="+model,
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run claude: %v", err)
	}

	return nil
}
