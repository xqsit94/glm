package cmd

import (
	"github.com/xqsit94/glm/internal/installer"

	"github.com/spf13/cobra"
)

func InstallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install tools",
		Long:  "Install various tools and dependencies",
	}

	cmd.AddCommand(installClaudeCmd())

	return cmd
}

func installClaudeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "claude",
		Short: "Install Claude Code",
		Long:  "Install Claude Code using npm",
		RunE: func(cmd *cobra.Command, args []string) error {
			return installer.InstallClaude()
		},
	}
}
