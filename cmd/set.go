package cmd

import (
	"fmt"

	"glm/internal/glm"

	"github.com/spf13/cobra"
)

func SetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Change GLM model settings",
		Long:  "Update the ANTHROPIC_MODEL in existing GLM configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			model, _ := cmd.Flags().GetString("model")

			if model == "" {
				return fmt.Errorf("model flag is required")
			}

			return glm.SetModel(model)
		},
	}

	cmd.Flags().StringP("model", "m", "", "GLM model to set (required)")
	cmd.MarkFlagRequired("model")

	return cmd
}
