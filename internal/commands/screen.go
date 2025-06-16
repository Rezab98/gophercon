package commands

import (
	"fmt"

	"github.com/Rezab98/gophercon/internal/automation"
	"github.com/spf13/cobra"
)

var screenCmd = &cobra.Command{
	Use:   "screen",
	Short: "Screen capture and analysis commands",
	Long:  "Commands for taking screenshots, finding images, and screen analysis",
}

var screenshotCmd = &cobra.Command{
	Use:   "capture [filename]",
	Short: "Take a screenshot",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		if len(args) > 0 {
			filename = args[0]
		} else {
			filename = "screenshot.png"
		}

		fullscreen, _ := cmd.Flags().GetBool("fullscreen")
		x, _ := cmd.Flags().GetInt("x")
		y, _ := cmd.Flags().GetInt("y")
		width, _ := cmd.Flags().GetInt("width")
		height, _ := cmd.Flags().GetInt("height")

		var err error
		if fullscreen {
			err = automation.CaptureFullScreen(filename)
		} else if width > 0 && height > 0 {
			err = automation.CaptureRegion(filename, x, y, width, height)
		} else {
			err = automation.CaptureFullScreen(filename)
		}

		checkError(err)
		fmt.Printf("Screenshot saved to: %s\n", filename)
	},
}

var findImageCmd = &cobra.Command{
	Use:   "find [image_path]",
	Short: "Find an image on screen",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		imagePath := args[0]
		tolerance, _ := cmd.Flags().GetFloat64("tolerance")

		x, y, found, err := automation.FindImageOnScreen(imagePath, tolerance)
		checkError(err)

		if found {
			fmt.Printf("Image found at coordinates: (%d, %d)\n", x, y)
		} else {
			fmt.Println("Image not found on screen")
		}
	},
}

var getPixelCmd = &cobra.Command{
	Use:   "pixel [x] [y]",
	Short: "Get pixel color at coordinates",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		x, y, err := automation.ParseCoordinates(args[0], args[1])
		if err != nil {
			checkError(fmt.Errorf("invalid coordinates: %v", err))
		}

		color, err := automation.GetPixelColor(x, y)
		checkError(err)

		fmt.Printf("Pixel color at (%d, %d): %s\n", x, y, color)
	},
}

func init() {
	// Screenshot flags
	screenshotCmd.Flags().BoolP("fullscreen", "f", true, "Capture fullscreen")
	screenshotCmd.Flags().IntP("x", "x", 0, "X coordinate for region capture")
	screenshotCmd.Flags().IntP("y", "y", 0, "Y coordinate for region capture")
	screenshotCmd.Flags().IntP("width", "w", 0, "Width for region capture")
	screenshotCmd.Flags().IntP("height", "h", 0, "Height for region capture")

	// Find image flags
	findImageCmd.Flags().Float64P("tolerance", "t", 0.8, "Image matching tolerance (0.0-1.0)")

	// Add subcommands
	screenCmd.AddCommand(screenshotCmd)
	screenCmd.AddCommand(findImageCmd)
	screenCmd.AddCommand(getPixelCmd)
}
