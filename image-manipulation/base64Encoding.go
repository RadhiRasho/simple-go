package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)


func Base64Encoding() {
	// Create a blank image 10x20 pixels
	myImage := image.NewRGBA(image.Rect(0,0,10,20))

	// In-memory buffer to store png image
	// before we base 64 encode it
	var buff bytes.Buffer

	// The buffer satisfies the Writer interface so we can use it with Encode
	// In previous examples we encoded to a file, this time to a temp buffer
	png.Encode(&buff, myImage)

	// Encode the bytes in the buffer to a base64 string
	encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	// You can embed it in an html doc with this string
	htmlImage := fmt.Sprintf("<img src=\"data:image/png;base64,%s\"/>", encodedString)

	fmt.Println(htmlImage)
}