package cmd

import (
	"github.com/rodrigorodrigo/glm/internal/token"

	"github.com/spf13/cobra"
)

func TokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "Manage authentication token",
		Long:  "Manage your Anthropic authentication token",
	}

	cmd.AddCommand(tokenSetCmd())
	cmd.AddCommand(tokenShowCmd())
	cmd.AddCommand(tokenClearCmd())

	return cmd
}

func tokenSetCmd() *cobra.Command {
	var tokenValue string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set authentication token",
		Long:  "Set your Anthropic authentication token interactively or via --token flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			if tokenValue != "" {
				return token.SetNonInteractive(tokenValue)
			}
			return token.Set()
		},
	}

	cmd.Flags().StringVar(&tokenValue, "token", "", "Authentication token (non-interactive)")

	return cmd
}

func tokenShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current token",
		Long:  "Display the current authentication token (masked)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return token.Show()
		},
	}
}

func tokenClearCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clear",
		Short: "Clear authentication token",
		Long:  "Remove the stored authentication token",
		RunE: func(cmd *cobra.Command, args []string) error {
			return token.Clear()
		},
	}
}
