package config

import (
	"path/filepath"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := New()
	if cfg.Dir() == "" {
		t.Error("Expected non-empty config dir")
	}
	if cfg.DBPath() == "" {
		t.Error("Expected non-empty DB path")
	}

	expectedDBPath := filepath.Join(cfg.Dir(), "history.db")
	if cfg.DBPath() != expectedDBPath {
		t.Errorf("Expected DB path %s, got %s", expectedDBPath, cfg.DBPath())
	}
}
