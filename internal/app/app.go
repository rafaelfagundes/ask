package app

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rafaelfagundes/ask/internal/config"
	"github.com/rafaelfagundes/ask/internal/gemini"
	"github.com/rafaelfagundes/ask/internal/history"
	"github.com/rafaelfagundes/ask/internal/osinfo"
)

type App struct {
	Config  *config.Config
	History *history.Store
	Gemini  *gemini.Client
	OSInfo  *osinfo.OSInfo
}

func New() (*App, error) {
	cfg := config.New()

	db, err := sql.Open("sqlite3", cfg.DBPath())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	historyStore, err := history.NewStore(db)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize history store: %w", err)
	}

	osInfo := osinfo.Get()

	geminiClient, err := gemini.NewClient(osInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Gemini client: %w", err)
	}

	return &App{
		Config:  cfg,
		History: historyStore,
		Gemini:  geminiClient,
		OSInfo:  osInfo,
	}, nil
}

func (a *App) Close() error {
	if err := a.History.Close(); err != nil {
		return err
	}
	return a.Gemini.Close()
}
