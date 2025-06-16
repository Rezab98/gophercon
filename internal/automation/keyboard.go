package automation

import (
	"time"

	"github.com/go-vgo/robotgo"
)

// TypeText types the specified text with optional delay between keystrokes
func TypeText(text string, delay int) error {
	if delay > 0 {
		for _, char := range text {
			robotgo.TypeStr(string(char))
			d := time.Duration(delay) * time.Millisecond
			time.Sleep(d)
		}
	} else {
		robotgo.TypeStr(text)
	}
	return nil
}

// SendKeys sends key combinations or single keys (e.g. "ctrl+c")
func SendKeys(keys string) error {
	return robotgo.KeyTap(keys)
}
