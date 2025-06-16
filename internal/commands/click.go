package commands

import (
	"fmt"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/spf13/cobra"
)

// clickCmd represents the standalone `click` command.
var clickCmd = &cobra.Command{
	Use:   "click [x] [y]",
	Short: "Click at a specific screen coordinate",
	Long: `Click at the given X,Y coordinate on the screen.

Examples:
  desktop-automation click 100 200
  desktop-automation click 640 480`,
	DisableFlagParsing: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("requires exactly 2 coordinates (x y)")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Parse and validate incoming coordinates.
		x, y, err := automation.ParseCoordinates(args[0], args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		if x < 0 || y < 0 {
			checkError(fmt.Errorf("coordinates must be non-negative; received (%d,%d)", x, y))
		}

		// Show current cursor position before moving.
		curX, curY := automation.GetPosition()
		fmt.Printf("Current mouse position: (%d, %d)\n", curX, curY)

		// Perform the click using our automation layer.
		if err := automation.Click(x, y); err != nil {
			checkError(err)
		}

		fmt.Printf("Successfully clicked at (%d, %d)\n", x, y)
	},
}

func init() {
	rootCmd.AddCommand(clickCmd)
}
