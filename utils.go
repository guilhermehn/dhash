package dhash

import (
	"encoding/hex"
	"image/color"
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
