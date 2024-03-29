package main

import (
	"bufio"
	"fmt"
	"global/utils"
	"os"
)

func BufferedReader() {
	path := "text-files/BufferedReader.txt"

	ExistsOrCreate(path)

	// Open file and create a buffered read on top
	file, err := os.Open(path)

	utils.FatalError(err)

	defer file.Close()

	bufferedReader := bufio.NewReader(file)

	// Get bytes without advancing pointer
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(len(byteSlice))

	utils.FatalError(err)

	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// Read and advance pointer
	numBytesRead, err := bufferedReader.Read(byteSlice)

	utils.FatalError(err)

	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// Read 1 byte. Error if no byte to read
	myByte, err := bufferedReader.ReadByte()

	utils.FatalError(err)

	fmt.Printf("Read 1 byte: %c\n", myByte)

	// Read up to and including delimiter
	// Returns byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')

	utils.FatalError(err)

	fmt.Printf("Read string: %s\n", dataBytes)

	// Read up to and including delimiter
	// Returns string

	dataString, err := bufferedReader.ReadString('\n')

	utils.FatalError(err)

	fmt.Printf("Read string: %s\n", dataString)
	// This example reads a few lines so test.txt
	// should have a few lines of text to work correctly
}
