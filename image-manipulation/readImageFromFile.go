package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func ReadImageFromFile() {
	// Read image from file that already exists
	existingImageFile, err := os.Open("WritingImageToFile.png")

	FatalError(err)

	defer existingImageFile.Close()

	// Calling the generic image.Decode() will give us the data
	// and type of image it is as a string. we expect "png"
	imageData, imageType, err := image.Decode(existingImageFile)

	FatalError(err)

	fmt.Println("Image Data: ", imageData)
	fmt.Println("\n\nImage Type: ", imageType)

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0,0)

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)

	FatalError(err)

	fmt.Println("LoadedImage: ",loadedImage)

}