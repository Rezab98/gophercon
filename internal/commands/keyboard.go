package commands

import (
	"fmt"
	"strings"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/spf13/cobra"
)

var keyboardCmd = &cobra.Command{
	Use:   "keyboard",
	Short: "Keyboard control commands",
	Long:  "Commands for typing text and sending key combinations",
}

var typeCmd = &cobra.Command{
	Use:   "type [text]",
	Short: "Type the specified text",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, " ")
		delay, _ := cmd.Flags().GetInt("delay")

		err := automation.TypeText(text, delay)
		checkError(err)

		fmt.Printf("Typed: %s\n", text)
	},
}

var keyCmd = &cobra.Command{
	Use:   "key [keys...]",
	Short: "Send key combination",
	Long: `Send single key or key combination.
Examples:
  desktop-automation keyboard key enter
  desktop-automation keyboard key ctrl+c
  desktop-automation keyboard key alt+tab`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keys := strings.Join(args, " ")

		err := automation.SendKeys(keys)
		checkError(err)

		fmt.Printf("Sent keys: %s\n", keys)
	},
}

func init() {
	// Type command flags
	typeCmd.Flags().IntP("delay", "d", 0, "Delay between keystrokes in milliseconds")

	// Add subcommands
	keyboardCmd.AddCommand(typeCmd)
	keyboardCmd.AddCommand(keyCmd)
}
