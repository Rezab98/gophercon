package commands

import (
	"fmt"
	"strings"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/Rezab98/gophercon/internal/ui"
	"github.com/spf13/cobra"
)

// typeTextCmd represents the standalone `type` command.
var typeTextCmd = &cobra.Command{
	Use:   "type [text]",
	Short: "Type text at the current cursor position",
	Long: `Type the provided text at the current cursor position.

Examples:
  desktop-automation type "Hello, World!"
  desktop-automation type The quick brown fox`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")

		// Validate and perform the typing action using the automation helpers.
		var err error
		if delayMS > 0 {
			err = automation.TypeStringDelay(text, delayMS)
		} else {
			err = automation.TypeString(text)
		}
		checkError(err)

		// Print success information.
		charCount := len([]rune(text))
		fmt.Println(ui.SuccessStyle.Render(fmt.Sprintf("Successfully typed %d characters.", charCount)))
	},
}

var (
	// delayMS holds the user-provided delay flag for the `type` command.
	delayMS int
)

func init() {
	// Register the flag before adding command to root
	typeTextCmd.Flags().IntVarP(&delayMS, "delay", "d", 0, "Delay between keystrokes in milliseconds")
	rootCmd.AddCommand(typeTextCmd)
}
