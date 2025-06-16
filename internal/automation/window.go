package automation

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// WindowInfo represents window information
type WindowInfo struct {
	Title string
	PID   int
}

// ListWindows returns a list of open windows
func ListWindows() ([]WindowInfo, error) {
	pids, err := robotgo.Pids()
	if err != nil {
		return nil, err
	}

	var windows []WindowInfo
	for _, pid := range pids {
		title := robotgo.GetTitle(pid)
		if title != "" {
			windows = append(windows, WindowInfo{
				Title: title,
				PID:   pid,
			})
		}
	}

	return windows, nil
}

func findPIDByTitle(title string) (int, error) {
	ids, err := robotgo.FindIds(title)
	if err != nil {
		return 0, err
	}
	if len(ids) == 0 {
		return 0, fmt.Errorf("window '%s' not found", title)
	}
	return ids[0], nil
}

// ActivateWindow activates a window by title
func ActivateWindow(title string) error {
	return robotgo.ActiveName(title)
}

// CloseWindow closes a window by title
func CloseWindow(title string, force bool) error {
	pid, err := findPIDByTitle(title)
	if err != nil {
		return err
	}
	if force {
		return robotgo.Kill(pid)
	}
	// Attempt graceful close
	robotgo.CloseWindow(pid)
	return nil
}

// MoveWindow moves a window to specified coordinates (not supported in current robotgo)
func MoveWindow(title string, x, y int) error {
	return fmt.Errorf("MoveWindow is not supported with current robotgo version")
}

// ResizeWindow resizes a window to specified dimensions (not supported in current robotgo)
func ResizeWindow(title string, width, height int) error {
	return fmt.Errorf("ResizeWindow is not supported with current robotgo version")
}
