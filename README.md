# Desktop Automation CLI

A Go-based desktop automation CLI tool for controlling mouse, keyboard, screen capture, and window management.

## Project Structure

```
├── cmd/
│   └── desktop-automation/
│       └── main.go              # Main entry point
├── internal/
│   ├── automation/              # Robot Go wrappers
│   │   ├── mouse.go
│   │   ├── keyboard.go
│   │   ├── screen.go
│   │   └── window.go
│   ├── commands/                # CLI commands
│   │   ├── root.go
│   │   ├── mouse.go
│   │   ├── keyboard.go
│   │   ├── screen.go
│   │   ├── window.go
│   │   └── tui.go
│   └── ui/                      # Bubble Tea TUI components
│       └── tui.go
├── go.mod
├── .gitignore
├── Taskfile.yml
└── README.md
```

## Usage

### Build and Run

```bash
task build
task run
```

### Commands

- `desktop-automation mouse` - Mouse control commands
- `desktop-automation keyboard` - Keyboard input commands  
- `desktop-automation screen` - Screen capture commands
- `desktop-automation window` - Window management commands
- `desktop-automation tui` - Launch TUI interface

### Development

```bash
task dev    # Run in development mode
task test   # Run tests
task clean  # Clean build artifacts
``` 