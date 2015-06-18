package main

import (
	"fmt"
	"github.com/disintegration/imaging"
	"os"
)

func main() {
	// Check if a file was passed
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [image]\n", os.Args[0])
		os.Exit(2)
	}

	// Read the image
	image, err := imaging.Open(os.Args[1])

	// Exit if could not read the image
	if err != nil {
		fmt.Println("Error loading the image: %s", err)
		os.Exit(1)
	}

	// Downscale to 9x8 and convert it to grayscale
	downscaledImage := imaging.Grayscale(imaging.Resize(image, 9, 8, imaging.Lanczos))

	fmt.Println(downscaledImage.Bounds())
}
