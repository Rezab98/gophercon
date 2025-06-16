package automation

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

// ClickMouse clicks the specified mouse button at the given coordinates
func ClickMouse(x, y int, button string) error {
	// Move to coordinates first using smooth movement
	robotgo.MoveSmooth(x, y, 1.0, 10.0)

	// Click the specified button
	switch button {
	case "left":
		robotgo.Click("left")
	case "right":
		robotgo.Click("right")
	case "middle":
		robotgo.Click("center")
	default:
		return fmt.Errorf("invalid mouse button: %s (use left, right, or middle)", button)
	}

	return nil
}

// MoveMouse moves the mouse cursor to the specified coordinates
func MoveMouse(x, y int) error {
	robotgo.MoveSmooth(x, y, 1.0, 10.0)
	return nil
}

// ScrollMouse scrolls in the specified direction by the given amount
func ScrollMouse(direction string, amount int) error {
	switch direction {
	case "up":
		robotgo.ScrollDir(amount, "up")
	case "down":
		robotgo.ScrollDir(amount, "down")
	case "left":
		robotgo.ScrollDir(amount, "left")
	case "right":
		robotgo.ScrollDir(amount, "right")
	default:
		return fmt.Errorf("invalid scroll direction: %s (use up, down, left, or right)", direction)
	}

	return nil
}

// GetMousePosition returns the current mouse position
func GetMousePosition() (int, int) {
	x, y := robotgo.Location()
	return x, y
}

// ParseCoordinates parses string coordinates to integers
func ParseCoordinates(xStr, yStr string) (int, int, error) {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid x coordinate: %s", xStr)
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid y coordinate: %s", yStr)
	}

	return x, y, nil
}

// ParseAmount parses string amount to integer
func ParseAmount(amountStr string) (int, error) {
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %s", amountStr)
	}

	return amount, nil
}

// Click is a convenience wrapper that validates the provided coordinates and performs a left-button click.
// It is intentionally simple because the higher-level CLI currently assumes a left-click. If you need to
// support additional buttons in the future, call ClickMouse directly.
func Click(x, y int) error {
	// Basic validation – coordinates must be zero or positive. Negative coordinates are not supported
	// because most window systems treat the origin (0,0) as the top-left corner of the primary screen.
	if x < 0 || y < 0 {
		return fmt.Errorf("coordinates must be non-negative; received (%d,%d)", x, y)
	}

	// Delegate to the existing, battle-tested helper.
	return ClickMouse(x, y, "left")
}

// GetPosition returns the current mouse cursor coordinates. This is a tiny wrapper around
// GetMousePosition to satisfy the naming convention required by the commands package.
func GetPosition() (int, int) {
	return GetMousePosition()
}

// Move moves the mouse cursor instantly (no built-in smoothing).
// This is useful for CLI commands that need immediate positioning without any visible animation.
func Move(x, y int) error {
	// Validate coordinates – negatives are not supported.
	if x < 0 || y < 0 {
		return fmt.Errorf("coordinates must be non-negative; received (%d,%d)", x, y)
	}

	// robotgo.MoveMouse performs an immediate jump without easing.
	robotgo.MoveMouse(x, y)
	return nil
}

// MoveSmooth moves the mouse cursor to the specified coordinates over the requested duration (in
// milliseconds). A very small duration effectively behaves like Move.  Internally this performs a
// linear interpolation between the current location and the target, sleeping between small steps
// so that the overall motion takes approximately the desired amount of time.
func MoveSmooth(x, y int, durationMs int) error {
	// Basic validation to keep the behaviour consistent with Move.
	if x < 0 || y < 0 {
		return fmt.Errorf("coordinates must be non-negative; received (%d,%d)", x, y)
	}
	if durationMs <= 0 {
		// A non-positive duration defaults to an instant move.
		return Move(x, y)
	}

	// Capture starting position so we can interpolate.
	startX, startY := robotgo.Location()
	dx := x - startX
	dy := y - startY

	// Number of steps.  We cap at 240 steps (~60fps over 4s) to avoid too many robotgo calls.
	steps := 120 // default granularity
	if durationMs > 4000 {
		steps = 240
	} else if durationMs < 500 {
		steps = 60
	}

	sleepPerStep := time.Duration(durationMs) * time.Millisecond / time.Duration(steps)

	for i := 1; i <= steps; i++ {
		// Linear interpolation.
		nx := startX + dx*i/steps
		ny := startY + dy*i/steps
		robotgo.MoveMouse(nx, ny)
		time.Sleep(sleepPerStep)
	}

	// Ensure we end exactly on target.
	robotgo.MoveMouse(x, y)
	return nil
}
