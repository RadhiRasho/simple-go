package main

import (
	"global/utils"
	"io"
	"log"
	"os"
)

func CopyFile() {
	// Copy a file
	// Open original file
	path := "text-files/copy.txt"

	ExistsOrCreate(path)

	original, err := os.Open(path)

	utils.FatalError(err)

	defer original.Close()

	// Create new copy
	newFile, err := os.Create("text-files/test_copy.txt")

	utils.FatalError(err)

	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, original)

	utils.FatalError(err)

	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file content
	// Flushes Memory To Disk
	err = newFile.Sync()
	utils.FatalError(err)
}