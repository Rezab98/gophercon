package automation

import (
	"fmt"
	"strconv"

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
