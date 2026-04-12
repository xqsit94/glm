package cmd

import (
	"fmt"

	"github.com/rodrigorodrigo/glm/internal/token"

	"github.com/spf13/cobra"
)

var availableModels = []string{
	"glm-5.1",
	"glm-5",
	"glm-4.7",
	"glm-4.6",
	"glm-4.5",
	"glm-4.5-air",
}

func ModelsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "models",
		Short: "List available GLM models",
		Long:  "Display all available GLM models with the current default marked",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Available GLM models:")
			fmt.Println()
			for _, model := range availableModels {
				if model == token.DefaultModel {
					fmt.Printf("  * %s (default)\n", model)
				} else {
					fmt.Printf("    %s\n", model)
				}
			}
		},
	}
}
