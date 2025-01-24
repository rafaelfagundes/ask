#!/bin/bash
set -e

# Install glow for markdown rendering
echo "Installing glow markdown viewer..."
go install github.com/charmbracelet/glow@latest

# Initialize and tidy module dependencies
echo "Setting up Go module..."
go mod init github.com/rafaelfagundes/ask >/dev/null 2>&1 || true
go mod tidy

# Build from the cmd/ask directory
echo "Building executable..."
go build -o ask ./cmd/ask

# Install to system binaries
echo "Installing system-wide (requires sudo)..."
sudo mv ask /usr/local/bin

echo "Installation complete! You can now run 'ask' from any directory."