# Ask 🤖 - Terminal AI Assistant powered by Gemini

[![Go Version](https://img.shields.io/github/go-mod/go-version/rafaelfagundes/ask)](https://golang.org/)
[![GitHub Release](https://img.shields.io/github/v/release/rafaelfagundes/ask)](https://github.com/rafaelfagundes/ask/releases)
[![Build and Release](https://github.com/rafaelfagundes/ask/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/rafaelfagundes/ask/actions/workflows/build.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A context-aware terminal assistant that leverages Google's Gemini AI to help with command line tasks, scripting, and system administration.

![Demo](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExd3I0ZXB3OWV6bXJ1a2k4Y2h4Y2U1M2J2b2p3dXJ2bnVkNmRjZ2N0diZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/26tn33aiTi1jkl6H6/giphy.gif)

## Features ✨

- 🗃️ Persistent command history with timestamp tracking
- 📖 Automatic markdown rendering with [glow](https://github.com/charmbracelet/glow)
- 💻 OS/shell context-aware responses
- 🗑️ Safe history deletion (single items or all)

## Installation ⚡

### Quick Install (Linux/macOS)
```bash
curl -sSL https://raw.githubusercontent.com/rafaelfagundes/ask/main/install.sh | bash
```

### Manual Installation
1. Install required dependencies:
```bash
go install github.com/charmbracelet/glow@latest
```
2. Build and install:
```bash
go install github.com/rafaelfagundes/ask/cmd/ask@latest
```

## Usage 🚀

### Basic Query
```bash
ask "How do I recursively find and delete .tmp files?"
```

### Interactive Mode
```bash
ask
> Enter your question: How to monitor CPU usage?
```

### History Management
```bash
ask history             # List all queries
ask history delete 3    # Delete item 3
ask history delete all  # Clear all history
ask last                # Show last response
```

## Configuration ⚙️

1. Get Gemini API Key from [Google AI Studio](https://aistudio.google.com/)
2. Add to your shell profile:
```bash
export GEMINI_API_KEY="your-api-key-here"
```

Configuration files stored in:  
`${XDG_CONFIG_HOME:-$HOME/.config}/ask`

## Examples 💡

**System Administration**  
`ask "Quick checklist for securing a new Ubuntu server"`

**Scripting Help**  
`ask "Create a bash script to backup directory with timestamp"`

**Troubleshooting**  
`ask "Debug 'Permission Denied' error after chmod"`

**DevOps**  
`ask "Explain Kubernetes pod lifecycle in simple terms"`

## Contributing 🤝

We welcome contributions! Please follow these steps:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

## Acknowledgments 🙏

- Google Gemini API for AI capabilities
- Charmbracelet Glow for markdown rendering
- SQLite for lightweight history storage

