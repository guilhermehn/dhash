package dhash

import (
	"github.com/disintegration/imaging"
	"strings"
)

// Runs a difference hash algorithm on the image
// and returns a hash string and a error if there
// was one during the function execution
func Dhash(path string, size int) (string, error) {
	// Read the image
	original, err := imaging.Open(path)

	// Return if could not read the image
	if err != nil {
		return "", err
	}

	// Downscale to (size+1)x(size) and convert it to grayscale
	img := imaging.Grayscale(imaging.Resize(original, size+1, size, imaging.Lanczos))

	// Create the intensity map
	intensities := make([]string, 0, size)

	// Iterate through the columns
	for y := 0; y < size; y++ {
		line := make([]byte, 0, size)

		// For each pixel of the line until the penultimate
		for x := 0; x < size; x++ {
			a := intensityOfColor(img.At(x, y))
			b := intensityOfColor(img.At(x+1, y))

			var diff byte

			// If the actual pixel is lighter than
			// the next one the diff should be 1
			if a > b {
				diff = 1
			}

			line = append(line, diff)
		}

		intensities = append(intensities, binarySliceToHex(line))
	}

	return strings.Join(intensities, ""), nil
}
