package osinfo

import (
	"os"
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	// Save original env vars
	origShell := os.Getenv("SHELL")
	origTerm := os.Getenv("TERM")
	defer func() {
		os.Setenv("SHELL", origShell)
		os.Setenv("TERM", origTerm)
	}()

	// Set test environment
	os.Setenv("SHELL", "/bin/bash")
	os.Setenv("TERM", "xterm")

	info := Get()

	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"OS", info.OS, runtime.GOOS},
		{"Shell", info.Shell, "/bin/bash"},
		{"Terminal", info.Terminal, "xterm"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("Get() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}

	// Version should not be empty on supported platforms
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" || runtime.GOOS == "windows" {
		if info.Version == "" {
			t.Error("Expected Version to be set")
		}
	}
}
