package main

import (
	"fmt"
	"os"
)




func SeekPositionInFile() {
	path := "text-files/seekPosition.txt"

	ExistsOrCreate(path)

	file, _ := os.Open(path)
	defer file.Close()

	// Offset is how many bytes to move
	// Offset can be positive or negative
	var offset int64 = 5

	// Whence is the point of reference for offset
	// 0 = Beginning of the file
	// 1 = current position
	// 2 = End of File
	var whence int = 0

	newPosition, err := file.Seek(offset, whence)

	FatalError(err)

	fmt.Println("Just moved to 5: ", newPosition)

	// Go back 2 bytes from current possition
	newPosition, err = file.Seek(-2, 1)

	FatalError(err)

	fmt.Println("Just moved back two: ", newPosition)

	// Find the current position by getting the
	// return value from seek after moving 0 bytes
	newPosition, err = file.Seek(0, 1)

	FatalError(err)

	fmt.Println("Current Position: ", newPosition)

	// Go To Beginning of file
	newPosition, err = file.Seek(0, 0)
	FatalError(err)

	fmt.Println("Position after seeking 0,0: ", newPosition)
}
