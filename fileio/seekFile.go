package main

import (
	"fmt"
	"global/utils"
	"os"
)

func SeekFile() {
	fmt.Println("Seeking out file")
	path := "text-files/seek.txt"
	// Simple read only open. We will cover actually reading
	// and writing to files in examples further down the page

	ExistsOrCreate(path)

	file, err := os.Open(path)

	utils.FatalError(err)

	fmt.Println("Close initial File Seek")
	file.Close()

	// OpenFile with more options. Last param is the permission mode
	// Second param is the attributes when opening
	fmt.Println("Secondary File seek, but with ")
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	utils.FatalError(err)

	file.Close()
	// Use these attributes individually or combined
	// with an OR for second arg of OpenFile()
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening
}