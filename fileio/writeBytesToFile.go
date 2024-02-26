package main

import (
	"log"
	"os"
)



func WriteBytesToFile() {
	path := "writeBytes.txt"

	ExistsOrCreate(path)

	// Open a new file for writing only
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	FatalError(err)

	defer file.Close()

	//Write bytes to file
	bytesSlice := []byte("Bytes!\n")

	bytesWritten, err := file.Write(bytesSlice)

	FatalError(err)

	log.Printf("Wrote %d bytes. \n", bytesWritten)
}
