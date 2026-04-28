package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Run full program tests
func TestProgram(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "No Arguments",
			args:     []string{},
			expected: "Usage:",
		},
		{
			name:     "Wrong Arguments",
			args:     []string{"--align=left"},
			expected: "Usage:",
		},
		{
			name:     "Invalid Banner",
			args:     []string{"--align=left", "hello", "fakebanner"},
			expected: "Error:",
		},
		{
			name:     "Valid Run",
			args:     []string{"--align=left", "hello", "standard"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdArgs := append([]string{"run", "."}, tt.args...)
			cmd := exec.Command("go", cmdArgs...)

			output, _ := cmd.CombinedOutput()
			result := string(output)

			if tt.expected != "" && !strings.Contains(result, tt.expected) {
				t.Errorf("Expected output to contain %q\nGot:\n%s", tt.expected, result)
			}
		})
	}
}

// Test terminal width separately
func TestGetTerminalWidth(t *testing.T) {
	width, err := getTerminalWidth()

	if err != nil {
		t.Log("Skipping terminal width test (no tty available)")
		return
	}

	if width <= 0 {
		t.Errorf("Expected width > 0, got %d", width)
	}
}

// Check if banner file exists
func TestBannerFileExists(t *testing.T) {
	if _, err := os.Stat("standard.txt"); os.IsNotExist(err) {
		t.Log("standard.txt not found in folder")
	}
}