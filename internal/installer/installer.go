package installer

import (
	"fmt"
	"os"
	"os/exec"
)

func InstallClaude() error {
	if !IsNpmAvailable() {
		fmt.Println("❌ npm is not available on your system.")
		fmt.Println("📦 To install Claude Code, you need Node.js and npm.")
		fmt.Println("🔗 Please install Node.js from: https://nodejs.org/")
		fmt.Println("💡 After installing Node.js, npm will be available automatically.")
		fmt.Println("🔄 Then run 'glm install claude' again.")
		return fmt.Errorf("npm not found")
	}

	fmt.Println("📦 Installing Claude Code...")
	fmt.Println("🔄 Running: npm install -g @anthropic-ai/claude-code")

	cmd := exec.Command("npm", "install", "-g", "@anthropic-ai/claude-code")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install Claude Code: %v", err)
	}

	fmt.Println("✅ Claude Code has been installed successfully!")
	fmt.Println("🚀 You can now use 'claude' command from anywhere.")
	return nil
}

func IsNpmAvailable() bool {
	_, err := exec.LookPath("npm")
	return err == nil
}
