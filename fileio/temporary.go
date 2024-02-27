package main

import (
	"bufio"
	"fmt"
	"global/utils"
	"log"
	"os"
	"time"
)


func TemporaryFilesAndDirectories() {
	// Create a temporary directory in the system default temp folder
    tempDirPath, err := os.MkdirTemp("", "MyTempDir") // creates MyTempDir in /tmp (on linux)

	utils.FatalError(err)

	fmt.Println("Temp Dir created: ", tempDirPath)

	// Create a file in new temporary directory
	tempFile, err := os.CreateTemp(tempDirPath, "TempFile.txt")

	utils.FatalError(err)

	fmt.Println("Temp File created: ", tempFile.Name())

	// ... Do something with the temporary file

	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(tempFile)

	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)

	utils.FatalError(err)

	log.Printf("bytes written: %d\n", bytesWritten)

	err = bufferedWriter.Flush()

	utils.FatalError(err)

	time.Sleep(10 * time.Second)

	// Close file
	err = tempFile.Close()

	utils.FatalError(err)

	// Delete the resources we created
	err = os.Remove(tempFile.Name())

	utils.FatalError(err)

	err = os.Remove(tempDirPath)

	utils.FatalError(err)
}