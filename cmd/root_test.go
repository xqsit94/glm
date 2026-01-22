package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestExtractUnknownFlags(t *testing.T) {
	// Create a test command with some known flags
	cmd := &cobra.Command{}
	cmd.Flags().StringP("model", "m", "", "Model to use")
	cmd.Flags().BoolP("yolo", "y", false, "Skip permissions")

	tests := []struct {
		name           string
		args           []string
		expectedUnknown []string
	}{
		{
			name:           "no args",
			args:           []string{},
			expectedUnknown: []string{},
		},
		{
			name:           "only known flags",
			args:           []string{"--model", "gpt-4"},
			expectedUnknown: []string{},
		},
		{
			name:           "unknown flag with value",
			args:           []string{"--allowedTools", "Bash,Read"},
			expectedUnknown: []string{"--allowedTools", "Bash,Read"},
		},
		{
			name:           "unknown flag with value and known flag",
			args:           []string{"--model", "gpt-4", "--allowedTools", "Bash,Read"},
			expectedUnknown: []string{"--allowedTools", "Bash,Read"},
		},
		{
			name:           "unknown flag with equals value",
			args:           []string{"--allowedTools=Bash,Read"},
			expectedUnknown: []string{"--allowedTools=Bash,Read"},
		},
		{
			name:           "unknown shorthand flag with value",
			args:           []string{"-x", "value"},
			expectedUnknown: []string{"-x", "value"},
		},
		{
			name:           "multiple unknown flags with values",
			args:           []string{"--foo", "bar", "--baz", "qux"},
			expectedUnknown: []string{"--foo", "bar", "--baz", "qux"},
		},
		{
			name:           "mix of known and unknown flags",
			args:           []string{"--model", "gpt-4", "--foo", "bar", "-y", "--baz", "qux"},
			expectedUnknown: []string{"--foo", "bar", "--baz", "qux"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original args and restore after test
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			// Set up os.Args for this test
			os.Args = append([]string{"glm"}, tt.args...)

			result := extractUnknownFlags(cmd)

			if !equalSlices(result, tt.expectedUnknown) {
				t.Errorf("extractUnknownFlags() = %v, want %v", result, tt.expectedUnknown)
			}
		})
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
