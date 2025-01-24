#!/bin/bash
set -e

REPO="github.com/rafaelfagundes/ask"
CLONE_DIR="$HOME/.ask-tmp"

# Clean up any previous failed installation
rm -rf "$CLONE_DIR"

# Clone the repository
echo "Cloning repository..."
git clone "https://$REPO.git" "$CLONE_DIR"
cd "$CLONE_DIR"

# Install glow for markdown rendering
echo "Installing glow markdown viewer..."
go install github.com/charmbracelet/glow@latest

# Initialize and tidy module dependencies
echo "Setting up Go module..."
if [ ! -f "go.mod" ]; then
    go mod init "$REPO"
fi

# Clean up module cache and rebuild
go clean -modcache
go mod tidy

# Build from the cmd/ask directory
echo "Building executable..."
go build -o ask ./cmd/ask

# Install to system binaries
echo "Installing system-wide (requires sudo)..."
sudo mv ask /usr/local/bin

# Cleanup
cd - > /dev/null
rm -rf "$CLONE_DIR"

echo "Installation complete! You can now run 'ask' from any directory."