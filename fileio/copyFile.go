package main

import (
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

	FatalError(err)

	defer original.Close()

	// Create new copy
	newFile, err := os.Create("text-files/test_copy.txt")

	FatalError(err)

	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, original)

	FatalError(err)

	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file content
	// Flushes Memory To Disk
	err = newFile.Sync()
	FatalError(err)
}