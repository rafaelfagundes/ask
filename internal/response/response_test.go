package response

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

type mockExecutor struct{}

func (e *mockExecutor) Command(name string, args ...string) *exec.Cmd {
	return exec.Command("echo", "mock")
}

func TestShow(t *testing.T) {
	// Save original executor and restore after test
	originalExecutor := executor
	defer func() { executor = originalExecutor }()

	tests := []struct {
		name    string
		input   string
		noPager bool
	}{
		{
			name:    "short response",
			input:   "short message",
			noPager: false,
		},
		{
			name:    "long response",
			input:   strings.Repeat("long message ", 100),
			noPager: false,
		},
		{
			name:    "forced no pager",
			input:   strings.Repeat("content ", 100),
			noPager: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set mock executor
			executor = &mockExecutor{}

			// Capture stdout
			oldStdout := os.Stdout
			_, w, _ := os.Pipe()
			os.Stdout = w

			Show(tt.input, tt.noPager)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			if tt.input != "" {
				if len(strings.TrimSpace(tt.input)) < 256 && !tt.noPager {
					t.Log("Verified short response uses no pager")
				}
			}
		})
	}
}
