package commands

import (
	"errors"
	"fmt"
	"strings"

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

		// TODO: Replace the stub below with the real type implementation.
		fmt.Println(ui.StubStyle.Render(fmt.Sprintf("[stub] Typing text: %s", text)))

		// Proper error handling placeholder
		checkError(errors.New("type command not implemented yet"))
	},
}

func init() {
	rootCmd.AddCommand(typeTextCmd)
}
