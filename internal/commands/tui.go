package commands

import (
	"fmt"

	"github.com/Rezab98/gophercon/internal/ui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the Terminal User Interface",
	Long:  "Launch an interactive TUI for desktop automation tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Launching Desktop Automation TUI...")

		err := ui.StartTUI()
		checkError(err)
	},
}
