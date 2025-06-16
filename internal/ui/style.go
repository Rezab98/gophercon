package ui

import "github.com/charmbracelet/lipgloss"

// Minimal styling definitions used across the CLI. Keeping this tiny on purpose.

var (
	// StubStyle colors stub messages to make them stand out in demos.
	StubStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Bold(true)

	// ErrorStyle colors error-like messages.
	ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)

	// SuccessStyle colors success messages.
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
)
