package app

import (
	"database/sql"

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
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", cfg.DatabasePath())
	if err != nil {
		return nil, err
	}

	historyStore, err := history.NewStore(db)
	if err != nil {
		db.Close()
		return nil, err
	}

	osInfo := osinfo.Get()

	geminiClient, err := gemini.NewClient(osInfo)
	if err != nil {
		return nil, err
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
