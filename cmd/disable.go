package cmd

import (
	"fmt"

	"glm/internal/glm"

	"github.com/spf13/cobra"
)

func DisableCmd() *cobra.Command {
	return &cobra.Command{
		Use:        "disable",
		Short:      "Disable GLM settings for Claude",
		Long:       "Remove GLM configuration and restore default Claude settings",
		Deprecated: "GLM now uses temporary session-based configuration. No need to disable - just run 'claude' directly.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("⚠️  Warning: This command is deprecated.")
			fmt.Println("💡 GLM now uses temporary session-based configuration.")
			fmt.Println("💡 To use Claude without GLM, just run 'claude' directly instead of 'glm'.")
			fmt.Println()

			return glm.Disable()
		},
	}
}
