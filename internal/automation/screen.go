package automation

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// CaptureFullScreen takes a full screen screenshot and saves it to filename
func CaptureFullScreen(filename string) error {
	img, err := robotgo.CaptureImg()
	if err != nil {
		return err
	}
	return robotgo.Save(img, filename)
}

// CaptureRegion captures a specific region of the screen and saves it to filename
func CaptureRegion(filename string, x, y, width, height int) error {
	img, err := robotgo.CaptureImg(x, y, width, height)
	if err != nil {
		return err
	}
	return robotgo.Save(img, filename)
}

// FindImageOnScreen finds an image on the current screen (simple template search)
func FindImageOnScreen(imagePath string, tolerance float64) (int, int, bool, error) {
	// robotgo provides FindImg API via cv sub-module, but for now we provide stub
	return 0, 0, false, fmt.Errorf("image search not implemented")
}

// GetPixelColor gets the color of a pixel at specified coordinates
func GetPixelColor(x, y int) (string, error) {
	color := robotgo.GetPixelColor(x, y)
	return fmt.Sprintf("#%s", color), nil
}
