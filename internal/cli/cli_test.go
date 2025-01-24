package cli

import (
	"testing"

	"github.com/rafaelfagundes/ask/internal/app"
)

func TestRun_NoArgs(t *testing.T) {
	a, err := app.New()
	if err != nil {
		t.Fatalf("App init failed: %v", err)
	}
	defer a.Close()

	if Run(a, []string{}) == nil {
		t.Log("Run succeeded with no args (expected).")
	}
}

func TestRun_HistoryCommand(t *testing.T) {
	a, err := app.New()
	if err != nil {
		t.Fatalf("App init failed: %v", err)
	}
	defer a.Close()

	if err := Run(a, []string{"history"}); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestRun(t *testing.T) {
	a, err := app.New()
	if err != nil {
		t.Fatalf("App init failed: %v", err)
	}
	defer a.Close()

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"last command", []string{"last"}, false},
		{"help flag", []string{"-h"}, false},
		{"help long flag", []string{"--help"}, false},
		{"config flag", []string{"-c"}, false},
		{"history command", []string{"history"}, false},
		{"history delete all", []string{"history", "delete", "all"}, false},
		{"invalid history position", []string{"history", "999999"}, true},
		{"invalid command", []string{"invalidcommand"}, false},
		{"no pager flag", []string{"--no-pager", "test question"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Run(a, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
