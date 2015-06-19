package dhash

import (
	"encoding/hex"
	"fmt"
	"github.com/disintegration/imaging"
	"image/color"
	"strings"
)

func getIntensityOfColor(c color.Color) uint8 {
	r, _, _, _ := c.RGBA()

	return uint8(r)
}

func binarySliceToHex(list []byte) string {
	var sum byte

	for _, n := range list {
		sum *= 2
		sum += n
	}

	return hex.EncodeToString([]byte{sum})
}

func Dhash(path string, size int) (string, error) {
	// Read the image
	image, err := imaging.Open(path)

	// Exit if could not read the image
	if err != nil {
		fmt.Println("Error loading the image: %s", err)
		return "", err
	}

	// Downscale to 9x8 and convert it to grayscale
	downscaledImage := imaging.Grayscale(imaging.Resize(image, size+1, size, imaging.Lanczos))

	bounds := downscaledImage.Bounds().Max
	x := bounds.X
	y := bounds.Y

	intensityMap := make([]string, 8)

	for i := 0; i < x; i++ {
		line := make([]byte, 8)

		for j := 0; j < y-1; j++ {
			actual := getIntensityOfColor(downscaledImage.At(i, j))
			next := getIntensityOfColor(downscaledImage.At(i, j+1))

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

	fmt.Println(strings.Join(intensityMap, ""))

	return "", nil
}
