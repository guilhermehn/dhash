package dhash

import (
	"encoding/hex"
	"github.com/disintegration/imaging"
	"image/color"
	"strings"
)

func intensityOfColor(c color.Color) uint32 {
	r, _, _, _ := c.RGBA()

	return r
}

func binarySliceToHex(slice []byte) string {
	var sum byte

	for _, value := range slice {
		sum *= 2
		sum += value
	}

	return hex.EncodeToString([]byte{sum})
}

func Dhash(path string, size int) (string, error) {
	// Read the image
	image, err := imaging.Open(path)

	// Exit if could not read the image
	if err != nil {
		return "", err
	}

	// Downscale to 9x8 and convert it to grayscale
	downscaledImage := imaging.Grayscale(imaging.Resize(image, size+1, size, imaging.Lanczos))

	// Get the image bounds (size)
	bounds := downscaledImage.Bounds().Max

	// Create the intensity map
	intensityMap := make([]string, 0, size)

	// Iterate through the lines
	for i := 0; i < bounds.X-1; i++ {
		line := make([]byte, 0, size)

		// For each pixel until the penultimate
		for j := 0; j < bounds.Y; j++ {
			actual := intensityOfColor(downscaledImage.At(i, j))
			next := intensityOfColor(downscaledImage.At(i+1, j))

			var difference byte

			if actual > next {
				difference = 1
			} else {
				difference = 0
			}

			line = append(line, difference)
		}

		intensityMap = append(intensityMap, binarySliceToHex(line))
	}

	return strings.Join(intensityMap, ""), nil
}
