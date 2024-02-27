package main

import (
	"global/utils"
	"image"
	"image/png"
	"math/rand"
	"os"
)

func WritingImageToFile() {
	// Create a blank image 100x200 pixels
	myImage := image.NewRGBA(image.Rect(0,0,100,200))

	for i := range (100*16) {
		myImage.Pix[i] = uint8(rand.Intn(255))
	}

	// output file is a file type which satisfies Writer interface
	outputFile, err := os.Create("WritingImageToFile.png")

	utils.FatalError(err)

	defer outputFile.Close()

	// Encode take a writer interface and an image interface
	// We pass it the File and the RGBA
	png.Encode(outputFile, myImage)

	// Don't forget to close files
	// Defered
}