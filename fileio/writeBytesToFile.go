package main

import (
	"global/utils"
	"log"
	"os"
)



func WriteBytesToFile() {
	path := "text-files/writeBytes.txt"

	ExistsOrCreate(path)

	// Open a new file for writing only
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	utils.FatalError(err)

	defer file.Close()

	//Write bytes to file
	bytesSlice := []byte("Bytes!\n")

	bytesWritten, err := file.Write(bytesSlice)

	utils.FatalError(err)

	log.Printf("Wrote %d bytes. \n", bytesWritten)
}
