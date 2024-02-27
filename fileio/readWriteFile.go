package main

import (
	"log"
	"os"
)


func ReadWriteFile() {
	// Test write permissions. It is possible the file
	// does not exit and that will return a different
	// error that can be checked with os.IsNotExist(err)
	path := "text-files/readWriteFile.txt"

	ExistsOrCreate(path)

	file, err := os.OpenFile(path, os.O_WRONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write Permission Denied.")
		}
	}

	file.Close()

	// Test read permissions
	file, err = os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil && os.IsPermission(err) {
		log.Println("Error: Read Permission Denied")
	}

	file.Close()
}
