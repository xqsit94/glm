package main

import (
	"fmt"
	"os"

	"github.com/rodrigorodrigo/glm/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	rootCmd.AddCommand(cmd.InstallCmd())
	rootCmd.AddCommand(cmd.TokenCmd())
	rootCmd.AddCommand(cmd.UpdateCmd())
	rootCmd.AddCommand(cmd.ModelsCmd())
	rootCmd.AddCommand(cmd.ConfigCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
