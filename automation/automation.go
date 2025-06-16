package automation

import (
	intauto "github.com/Rezab98/gophercon/internal/automation"
)

// Click performs a left-button click at the given coordinates.
// It is a thin wrapper around the internal automation package so that other
// modules in this monorepo can access mouse actions without violating the
// Go 'internal' visibility rule.
func Click(x, y int) error {
	return intauto.ClickMouse(x, y, "left")
}

// Move moves the mouse instantly to (x,y).
func Move(x, y int) error {
	return intauto.Move(x, y)
}

// MoveSmooth moves mouse to (x,y) over durationMs milliseconds.
func MoveSmooth(x, y int, durationMs int) error {
	return intauto.MoveSmooth(x, y, durationMs)
}

// Type types the provided text. If delayMs>0, it types with that delay per rune.
func Type(text string, delayMs int) error {
	if delayMs > 0 {
		return intauto.TypeStringDelay(text, delayMs)
	}
	return intauto.TypeString(text)
}
