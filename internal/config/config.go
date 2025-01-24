package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	dir string
}

func New() *Config {
	cfgDir, _ := os.UserConfigDir()
	return &Config{
		dir: filepath.Join(cfgDir, "ask"),
	}
}

func (c *Config) Dir() string {
	return c.dir
}

func (c *Config) DBPath() string {
	return filepath.Join(c.dir, "history.db")
}
