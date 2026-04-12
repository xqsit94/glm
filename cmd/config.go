package cmd

import (
	"fmt"

	"github.com/rodrigorodrigo/glm/internal/config"
	"github.com/rodrigorodrigo/glm/internal/token"

	"github.com/spf13/cobra"
)

func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage GLM configuration",
		Long:  "View and modify GLM configuration settings",
	}

	cmd.AddCommand(configShowCmd())
	cmd.AddCommand(configSetCmd())
	cmd.AddCommand(configResetCmd())

	return cmd
}

func configShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Show current configuration",
		Long:  "Display the current GLM configuration settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			model := cfg.DefaultModel
			if model == "" {
				model = token.DefaultModel
			}

			baseURL := cfg.BaseURL
			if baseURL == "" {
				baseURL = config.DefaultBaseURL
			}

			hasToken := cfg.AnthropicAuthToken != ""

			fmt.Println("GLM Configuration:")
			fmt.Println()
			fmt.Printf("  Model:    %s\n", model)
			fmt.Printf("  Base URL: %s\n", baseURL)
			fmt.Printf("  Token:    %s\n", formatTokenStatus(hasToken))

			return nil
		},
	}
}

func formatTokenStatus(hasToken bool) string {
	if hasToken {
		return "configured"
	}
	return "not set"
}

func configSetCmd() *cobra.Command {
	var model string
	var baseURL string
	var tokenValue string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set configuration values",
		Long:  "Set GLM configuration values like default model, base URL, or token",
		RunE: func(cmd *cobra.Command, args []string) error {
			if model == "" && baseURL == "" && tokenValue == "" {
				return fmt.Errorf("specify at least one of --model, --base-url, or --token")
			}

			if tokenValue != "" {
				if err := token.SetNonInteractive(tokenValue); err != nil {
					return err
				}
			}

			if model != "" || baseURL != "" {
				cfg, err := config.Load()
				if err != nil {
					return err
				}

				if model != "" {
					cfg.DefaultModel = model
					fmt.Printf("✅ Default model set to: %s\n", model)
				}

				if baseURL != "" {
					cfg.BaseURL = baseURL
					fmt.Printf("✅ Base URL set to: %s\n", baseURL)
				}

				if err := config.Save(cfg); err != nil {
					return err
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&model, "model", "", "Default GLM model")
	cmd.Flags().StringVar(&baseURL, "base-url", "", "API base URL")
	cmd.Flags().StringVar(&tokenValue, "token", "", "Authentication token")

	return cmd
}

func configResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Reset configuration to defaults",
		Long:  "Reset all GLM configuration settings to their default values",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return err
			}

			cfg.DefaultModel = ""
			cfg.BaseURL = ""

			if err := config.Save(cfg); err != nil {
				return err
			}

			fmt.Println("✅ Configuration has been reset to defaults.")
			fmt.Printf("   Model:    %s\n", token.DefaultModel)
			fmt.Printf("   Base URL: %s\n", config.DefaultBaseURL)

			return nil
		},
	}
}
