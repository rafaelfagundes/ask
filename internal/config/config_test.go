package config

import (
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg, err := New()
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}

	if cfg.Dir() == "" {
		t.Error("Expected non-empty config dir")
	}
	if cfg.DatabasePath() == "" {
		t.Error("Expected non-empty DB path")
	}

	expectedDBPath := filepath.Join(cfg.Dir(), "history.db")
	if cfg.DatabasePath() != expectedDBPath {
		t.Errorf("Expected DB path %s, got %s", expectedDBPath, cfg.DatabasePath())
	}
}
