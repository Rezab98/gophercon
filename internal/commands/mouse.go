package commands

import (
	"fmt"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/spf13/cobra"
)

var mouseCmd = &cobra.Command{
	Use:   "mouse",
	Short: "Mouse control commands",
	Long:  "Commands for controlling mouse movement, clicks, and scrolling",
}

var mouseClickCmd = &cobra.Command{
	Use:   "click [x] [y]",
	Short: "Click at specified coordinates",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		button, _ := cmd.Flags().GetString("button")

		x, y, err := automation.ParseCoordinates(args[0], args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		err = automation.ClickMouse(x, y, button)
		checkError(err)

		fmt.Printf("Clicked %s button at (%d, %d)\n", button, x, y)
	},
}

var mouseMoveCmd = &cobra.Command{
	Use:   "move [x] [y]",
	Short: "Move mouse to specified coordinates",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		x, y, err := automation.ParseCoordinates(args[0], args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		err = automation.MoveMouse(x, y)
		checkError(err)

		fmt.Printf("Moved mouse to (%d, %d)\n", x, y)
	},
}

var mouseScrollCmd = &cobra.Command{
	Use:   "scroll [direction] [amount]",
	Short: "Scroll in specified direction",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		direction := args[0]
		amount, err := automation.ParseAmount(args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid amount: %v", err))
		}

		err = automation.ScrollMouse(direction, amount)
		checkError(err)

		fmt.Printf("Scrolled %s by %d\n", direction, amount)
	},
}

func init() {
	// Mouse click flags
	mouseClickCmd.Flags().StringP("button", "b", "left", "Mouse button (left, right, middle)")

	// Add subcommands
	mouseCmd.AddCommand(mouseClickCmd)
	mouseCmd.AddCommand(mouseMoveCmd)
	mouseCmd.AddCommand(mouseScrollCmd)
}
