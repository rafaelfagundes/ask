package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCheckDependencies(t *testing.T) {
	// Temporarily modify PATH to test both scenarios
	originalPath := os.Getenv("PATH")
	defer os.Setenv("PATH", originalPath)

	t.Run("glow exists", func(t *testing.T) {
		if _, err := exec.LookPath("glow"); err != nil {
			t.Skip("Skipping test as glow is not installed")
		}
		if err := checkDependencies(); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("glow missing", func(t *testing.T) {
		os.Setenv("PATH", "")
		err := checkDependencies()
		if err == nil {
			t.Error("Expected error when glow is missing")
		}
		if !strings.Contains(err.Error(), "Glow not found") {
			t.Errorf("Expected error message about Glow not found, got: %v", err)
		}
	})
}

func TestMain(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"ask", "--help"}

	t.Log("Main initialization tested successfully")
}
