package main

import (
	"fmt"
	"os"

	"glm/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	rootCmd.AddCommand(cmd.EnableCmd())
	rootCmd.AddCommand(cmd.DisableCmd())
	rootCmd.AddCommand(cmd.SetCmd())
	rootCmd.AddCommand(cmd.InstallCmd())
	rootCmd.AddCommand(cmd.TokenCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
