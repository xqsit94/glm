package cmd

import (
	"fmt"

	"glm/internal/glm"
	"glm/internal/token"

	"github.com/spf13/cobra"
)

func EnableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:        "enable",
		Short:      "Enable GLM settings for Claude",
		Long:       "Configure Claude to use GLM model via BigModel API",
		Deprecated: "GLM now uses temporary session-based configuration. Just run 'glm' to launch Claude with GLM.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("⚠️  Warning: This command is deprecated.")
			fmt.Println("💡 Just run 'glm' to launch Claude with GLM using temporary configuration.")
			fmt.Println()

			model, _ := cmd.Flags().GetString("model")

			authToken, err := token.Get()
			if err != nil {
				return err
			}

			return glm.Enable(model, authToken)
		},
	}

	cmd.Flags().StringP("model", "m", defaultModel, fmt.Sprintf("GLM model to use (default: %s)", defaultModel))

	return cmd
}
