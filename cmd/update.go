package cmd

import (
	"fmt"
	"strings"

	"github.com/rodrigorodrigo/glm/internal/updater"

	"github.com/spf13/cobra"
)

func UpdateCmd() *cobra.Command {
	var checkOnly bool
	var force bool

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update GLM to the latest version",
		Long:  "Check for updates and install the latest version of GLM from GitHub",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpdate(checkOnly, force)
		},
	}

	cmd.Flags().BoolVar(&checkOnly, "check", false, "Only check for updates without installing")
	cmd.Flags().BoolVar(&force, "force", false, "Update without confirmation prompt")

	return cmd
}

func runUpdate(checkOnly, force bool) error {
	fmt.Println("🔍 Checking for updates...")
	fmt.Printf("📌 Current version: %s\n", version)

	info, err := updater.CheckForUpdate(version)
	if err != nil {
		fmt.Println("❌ Unable to check for updates. Please check your internet connection.")
		return fmt.Errorf("update check failed: %v", err)
	}

	if !info.HasUpdate {
		fmt.Println("✅ You're already running the latest version!")
		return nil
	}

	fmt.Printf("✨ Latest version: %s available!\n\n", info.LatestVersion)

	releaseNotes := updater.FormatReleaseNotes(info.ReleaseNotes, 10)
	if releaseNotes != "" {
		fmt.Println("📝 What's new:")
		for _, line := range strings.Split(releaseNotes, "\n") {
			if strings.TrimSpace(line) != "" {
				fmt.Printf("   %s\n", line)
			}
		}
		fmt.Println()
	}

	fmt.Printf("🔗 View full release notes: %s\n\n", info.ReleaseURL)

	if checkOnly {
		fmt.Printf("💡 Run 'glm update' to install version %s\n", info.LatestVersion)
		return nil
	}

	if !force {
		fmt.Printf("Would you like to update to %s? (y/N): ", info.LatestVersion)
		var response string
		fmt.Scanln(&response)

		response = strings.ToLower(strings.TrimSpace(response))
		if response != "y" && response != "yes" {
			fmt.Println("Update cancelled.")
			return nil
		}
	}

	osName, arch, err := updater.DetectPlatform()
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return err
	}

	fmt.Printf("\n📥 Downloading glm %s for %s/%s...\n", info.LatestVersion, osName, arch)

	var lastPercent int
	progressCallback := func(downloaded, total int64) {
		if total > 0 {
			percent := int(float64(downloaded) / float64(total) * 100)
			if percent > lastPercent {
				lastPercent = percent
				showProgress(percent, downloaded, total)
			}
		}
	}

	binaryPath, err := updater.DownloadBinary(info.LatestVersion, osName, arch, progressCallback)
	if err != nil {
		fmt.Printf("\n❌ Failed to download update: %v\n", err)
		fmt.Println("💡 Try again later or download manually from:")
		fmt.Printf("   %s\n", info.ReleaseURL)
		return err
	}

	fmt.Println("\n✅ Download complete!")

	fmt.Println("🔍 Verifying checksum...")
	if err := updater.VerifyChecksum(binaryPath, info.LatestVersion, osName, arch); err != nil {
		fmt.Printf("❌ Checksum verification failed: %v\n", err)
		return err
	}

	fmt.Println("🔧 Installing update...")

	if err := updater.VerifyBinary(binaryPath); err != nil {
		fmt.Printf("❌ Failed to verify downloaded binary: %v\n", err)
		return err
	}

	if err := updater.InstallUpdate(binaryPath); err != nil {
		fmt.Printf("❌ Failed to install update: %v\n", err)
		if strings.Contains(err.Error(), "permission denied") {
			fmt.Println("💡 Try running with sudo:")
			fmt.Println("   sudo glm update")
		}
		return err
	}

	fmt.Printf("✅ Successfully updated to %s!\n\n", info.LatestVersion)
	fmt.Println("🎉 GLM has been updated! The new version is now active.")

	return nil
}

func showProgress(percent int, downloaded, total int64) {
	barWidth := 40
	filled := barWidth * percent / 100
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	mb := float64(downloaded) / 1024 / 1024
	totalMB := float64(total) / 1024 / 1024

	fmt.Printf("\r[%s] %3d%% (%.1f/%.1f MB)", bar, percent, mb, totalMB)
}
