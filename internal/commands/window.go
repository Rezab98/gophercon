package commands

import (
	"fmt"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/spf13/cobra"
)

var windowCmd = &cobra.Command{
	Use:   "window",
	Short: "Window management commands",
	Long:  "Commands for managing application windows",
}

var listWindowsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all open windows",
	Run: func(cmd *cobra.Command, args []string) {
		windows, err := automation.ListWindows()
		checkError(err)

		fmt.Println("Open Windows:")
		for i, window := range windows {
			fmt.Printf("%d. %s (PID: %d)\n", i+1, window.Title, window.PID)
		}
	},
}

var activateWindowCmd = &cobra.Command{
	Use:   "activate [window_title]",
	Short: "Activate window by title",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		windowTitle := args[0]

		err := automation.ActivateWindow(windowTitle)
		checkError(err)

		fmt.Printf("Activated window: %s\n", windowTitle)
	},
}

var closeWindowCmd = &cobra.Command{
	Use:   "close [window_title]",
	Short: "Close window by title",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		windowTitle := args[0]
		force, _ := cmd.Flags().GetBool("force")

		err := automation.CloseWindow(windowTitle, force)
		checkError(err)

		fmt.Printf("Closed window: %s\n", windowTitle)
	},
}

var moveWindowCmd = &cobra.Command{
	Use:   "move [window_title] [x] [y]",
	Short: "Move window to specified coordinates",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		windowTitle := args[0]
		x, y, err := automation.ParseCoordinates(args[1], args[2])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		err = automation.MoveWindow(windowTitle, x, y)
		checkError(err)

		fmt.Printf("Moved window '%s' to (%d, %d)\n", windowTitle, x, y)
	},
}

var resizeWindowCmd = &cobra.Command{
	Use:   "resize [window_title] [width] [height]",
	Short: "Resize window to specified dimensions",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		windowTitle := args[0]
		width, height, err := automation.ParseCoordinates(args[1], args[2])
		if err != nil {
			checkError(fmt.Errorf("invalid dimensions: %v", err))
		}

		err = automation.ResizeWindow(windowTitle, width, height)
		checkError(err)

		fmt.Printf("Resized window '%s' to %dx%d\n", windowTitle, width, height)
	},
}

func init() {
	// Close window flags
	closeWindowCmd.Flags().BoolP("force", "f", false, "Force close the window")

	// Add subcommands
	windowCmd.AddCommand(listWindowsCmd)
	windowCmd.AddCommand(activateWindowCmd)
	windowCmd.AddCommand(closeWindowCmd)
	windowCmd.AddCommand(moveWindowCmd)
	windowCmd.AddCommand(resizeWindowCmd)
}
