package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
	"github.com/xqsit94/glm/internal/token"

	"github.com/spf13/cobra"
)

const (
	version = "1.2.0"
)

func RootCmd() *cobra.Command {
	var model string
	var yolo bool

	cmd := &cobra.Command{
		Use:     "glm",
		Short:   "GLM Claude settings management CLI",
		Long:    "A CLI tool to launch Claude with GLM settings using temporary session-based configuration",
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDefaultAction(cmd, model, yolo)
		},
	}

	cmd.Flags().StringVarP(&model, "model", "m", token.DefaultModel, "GLM model to use for this session")
	cmd.Flags().BoolVar(&yolo, "yolo", false, "Skip permission prompts (--dangerously-skip-permissions)")
	cmd.FParseErrWhitelist.UnknownFlags = true

	return cmd
}

// extractUnknownFlags extracts flags from os.Args that are not known to glm
// This allows passthrough of arbitrary flags to the claude command
func extractUnknownFlags(cmd *cobra.Command) []string {
	knownFlags := make(map[string]bool)
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		knownFlags["--"+f.Name] = true
		if f.Shorthand != "" {
			knownFlags["-"+f.Shorthand] = true
		}
	})

	var unknown []string
	args := os.Args[1:]

	skipNext := false
	for i, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}

		if knownFlags[arg] {
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				skipNext = true
			}
			continue
		}

		if strings.HasPrefix(arg, "--") {
			if idx := strings.Index(arg, "="); idx != -1 {
				flagName := arg[:idx]
				if !knownFlags[flagName] {
					unknown = append(unknown, arg)
				}
				continue
			}
		}

		if strings.HasPrefix(arg, "-") {
			unknown = append(unknown, arg)
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				unknown = append(unknown, args[i+1])
				skipNext = true
			}
		}
	}

	return unknown
}

func runDefaultAction(cmd *cobra.Command, model string, yolo bool) error {
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

	cmdArgs := []string{"claude", "--model", model}
	if yolo {
		cmdArgs = append(cmdArgs, "--dangerously-skip-permissions")
	}
	unknownFlags := extractUnknownFlags(cmd)
	cmdArgs = append(cmdArgs, unknownFlags...)

	claudeCmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	claudeCmd.Stdin = os.Stdin
	claudeCmd.Stdout = os.Stdout
	claudeCmd.Stderr = os.Stderr
	claudeCmd.Env = append(os.Environ(),
		"ANTHROPIC_BASE_URL=https://open.bigmodel.cn/api/anthropic",
		"ANTHROPIC_AUTH_TOKEN="+authToken,
		"ANTHROPIC_MODEL="+model,
	)

	if err := claudeCmd.Run(); err != nil {
		return fmt.Errorf("failed to run claude: %v", err)
	}

	return nil
}
