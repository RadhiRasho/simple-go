package main

import (
	"fmt"
	"global/utils"
	"io"
	"log"
	"os"
)

func ReadExactlyNBytes() {
	path := "text-files/readExactlyNBytes.txt"

	ExistsOrCreate(path)

	//Open File
	file, err := os.Open(path)

	utils.FatalError(err)

	defer file.Close()

	// The file.Read() function will happily read a tiny file in to a large
	// byte slice, but io.ReadFull() will return an
	// error if the file is smaller than the byte slice
	bytesSlice := make([]byte, 2) // 2 is the number of bytes that will be read
	numBytesRead, err := io.ReadFull(file, bytesSlice)

	utils.FatalError(err)

	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", bytesSlice)
}

func ReadUpNBytes() {
	path := "text-files/readUpNBytes.txt"

	ExistsOrCreate(path)
	// Open file for reading

	file, err := os.Open(path)

	utils.FatalError(err)

	defer file.Close()

	// Read up to len(b) bytes from the File
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	bytesSlice := make([]byte, 16) // will read up to the 16th byte within the file
	bytesRead, err := file.Read(bytesSlice)

	utils.FatalError(err)

	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", bytesSlice)
}

func ReadAtLeastNBytes() {
	path := "text-files/ReadAtLeastNBytes.txt"

	ExistsOrCreate(path)

	// Open file for reading
	file, err := os.Open(path)

	utils.FatalError(err)

	byteSlice := make([]byte, 512)

	minBytes := 8

	// io.ReadAtLeast() will return an error if it cannot
	// find at least minBytes to read. It will read as
	// many bytes as byteSlice can hold.
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)

	utils.FatalError(err)

	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

func ReadAllBytesOfFile() {
	path := "text-files/ReadAllBytesOfFile.txt"

	ExistsOrCreate(path)

	// Open file for reading
	file, err := os.Open(path)

	utils.FatalError(err)

	defer file.Close()

	// os.File.Read(), io.ReadFull(), and
	// io.ReadAtLeast() all work with a fixed
	// byte slice that you make before you read

	// io.ReadAll() will read every byte
	// from the read (in this case a file)
	// and return a slice of unknown slice
	data, err := io.ReadAll(file)

	utils.FatalError(err)

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d\n", len(data))
}


func QuickReadFileIntoMemory() {
	path := "text-files/QuickReadFileIntoMemory.txt"

	ExistsOrCreate(path)

	data, err := os.ReadFile(path)

	utils.FatalError(err)

	log.Printf("Data read: %s\n", data)
}
