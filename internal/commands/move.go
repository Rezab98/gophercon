package commands

import (
	"errors"
	"fmt"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/Rezab98/gophercon/internal/ui"
	"github.com/spf13/cobra"
)

// moveCmd represents the standalone `move` command.
var moveCmd = &cobra.Command{
	Use:   "move [x] [y]",
	Short: "Move the mouse cursor to coordinates",
	Long: `Move the mouse cursor to the specified X,Y coordinates.

Examples:
  desktop-automation move 500 300`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		x, y, err := automation.ParseCoordinates(args[0], args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		// TODO: Replace the stub below with the real move implementation.
		fmt.Println(ui.StubStyle.Render(fmt.Sprintf("[stub] Move cursor to (%d, %d)", x, y)))

		// Proper error handling placeholder
		checkError(errors.New("move command not implemented yet"))
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
