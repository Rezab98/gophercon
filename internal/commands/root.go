package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "desktop-automation",
	Short: "Beautiful Desktop Automation CLI",
	Long: `Beautiful Desktop Automation is a friendly command-line interface for performing common desktop automation tasks such as clicking, typing and moving the mouse. 
It is designed with clarity, helpful usage examples and sensible error handling so you can extend it without touching existing code.`,
	Version: "v0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		// If no subcommand is provided, show help
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet mode")

	// New standalone command stubs
	// NOTE: Legacy nested sub-trees (mouse, keyboard, etc.) were deliberately detached
	// because they conflict with the simplified CLI requested.

	// Subcommands are registered in their respective init functions to avoid duplication.
}

// checkError is a helper function to handle errors consistently
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
