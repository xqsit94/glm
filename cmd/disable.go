package cmd

import (
	"glm/internal/glm"

	"github.com/spf13/cobra"
)

func DisableCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disable",
		Short: "Disable GLM settings for Claude",
		Long:  "Remove GLM configuration and restore default Claude settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			return glm.Disable()
		},
	}
}
