package commands

import (
	"fmt"
	"math"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/spf13/cobra"
)

var (
	smoothFlag   bool
	durationFlag float64
)

// moveCmd represents the standalone `move` command.
var moveCmd = &cobra.Command{
	Use:   "move [x] [y]",
	Short: "Move the mouse cursor to coordinates",
	Long: `Move the mouse cursor to the specified X,Y coordinates.

Examples:
  desktop-automation move 500 300
  desktop-automation move --smooth --duration 5.0 100 100`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Parse CLI args into integers.
		x, y, err := automation.ParseCoordinates(args[0], args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		// Retrieve flag values.
		smoothFlag, _ = cmd.Flags().GetBool("smooth")
		durationFlag, _ = cmd.Flags().GetFloat64("duration")

		// Clamp negative duration values to zero.
		if durationFlag < 0 {
			durationFlag = 0
		}

		// Display current and target positions.
		curX, curY := automation.GetPosition()
		fmt.Printf("Current position: (%d, %d)\n", curX, curY)
		fmt.Printf("Target  position: (%d, %d)\n", x, y)

		// Display moving indicator.
		fmt.Println("Moving...")

		if smoothFlag {
			// Convert seconds -> milliseconds (rounded up)
			ms := int(math.Ceil(durationFlag * 1000))
			err = automation.MoveSmooth(x, y, ms)
		} else {
			err = automation.Move(x, y)
		}
		checkError(err)

		// Confirm final position.
		finalX, finalY := automation.GetPosition()
		fmt.Printf("Reached final position: (%d, %d)\n", finalX, finalY)
	},
}

func init() {
	// Define flags.
	moveCmd.Flags().BoolVarP(&smoothFlag, "smooth", "s", false, "Enable smooth movement animation")
	moveCmd.Flags().Float64VarP(&durationFlag, "duration", "d", 1.0, "Duration in seconds for smooth movement (only relevant with --smooth)")

	rootCmd.AddCommand(moveCmd)
}
