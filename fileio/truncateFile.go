package main

import (
	"fmt"
	"os"
)

func TruncateFile() {
	// Truncate a file to 100 bytes. If file
	// is less than 100 bytes the original contents will remain
	// at the beginning, and the rest of the space is
	// filled will null bytes. If it is over 100 bytes,
	// Everything past 100 bytes will be lost. Either way
	// we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely empty file
	fmt.Println("Truncating file")

	err := os.Truncate("text-files/truncation.txt", 100)

	FatalError(err)

	fmt.Println("File truncated")
}