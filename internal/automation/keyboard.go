package automation

import (
	"errors"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	// ErrEmptyText is returned when the provided text is empty or whitespace.
	ErrEmptyText = errors.New("text to type cannot be empty")
	// ErrNegativeDelay is returned when a negative delay is supplied.
	ErrNegativeDelay = errors.New("delay cannot be negative")
)

// TypeString types the whole text at once. It validates the text before sending.
func TypeString(text string) error {
	if strings.TrimSpace(text) == "" {
		return ErrEmptyText
	}
	robotgo.TypeStr(text)
	return nil
}

// TypeStringDelay types the provided text rune-by-rune, waiting delayMS milliseconds between strokes.
func TypeStringDelay(text string, delayMS int) error {
	if strings.TrimSpace(text) == "" {
		return ErrEmptyText
	}
	if delayMS < 0 {
		return ErrNegativeDelay
	}
	if delayMS == 0 {
		return TypeString(text)
	}
	d := time.Duration(delayMS) * time.Millisecond
	for _, char := range text {
		robotgo.TypeStr(string(char))
		time.Sleep(d)
	}
	return nil
}

// TypeText is kept for backwards compatibility. It delegates to the newer helpers.
func TypeText(text string, delay int) error {
	if delay > 0 {
		return TypeStringDelay(text, delay)
	}
	return TypeString(text)
}

// SendKeys sends key combinations or single keys (e.g. "ctrl+c")
func SendKeys(keys string) error {
	return robotgo.KeyTap(keys)
}
