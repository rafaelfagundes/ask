package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	dir string
}

func New() (*Config, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	return &Config{dir: configDir}, nil
}

func (c *Config) Dir() string {
	return c.dir
}

func (c *Config) DatabasePath() string {
	return filepath.Join(c.dir, "history.db")
}

func getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".ask"), nil
}
