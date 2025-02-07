package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rafaelfagundes/ask/internal/app"
	"github.com/rafaelfagundes/ask/internal/cli"
)

func checkDependencies() error {
	if _, err := exec.LookPath("glow"); err != nil {
		return fmt.Errorf("glow not found. Install with:\n  go install github.com/charmbracelet/glow@latest")
	}
	return nil
}

func main() {
	if err := checkDependencies(); err != nil {
		log.Fatal(err)
	}

	askApp, err := app.New()
	if err != nil {
		log.Fatal("Failed to initialize application:", err)
	}
	defer askApp.Close()

	if err := cli.Run(askApp, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
