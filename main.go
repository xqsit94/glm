package main

import (
	"fmt"
	"os"

	"github.com/xqsit94/glm/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	rootCmd.AddCommand(cmd.EnableCmd())
	rootCmd.AddCommand(cmd.DisableCmd())
	rootCmd.AddCommand(cmd.InstallCmd())
	rootCmd.AddCommand(cmd.TokenCmd())
	rootCmd.AddCommand(cmd.UpdateCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
