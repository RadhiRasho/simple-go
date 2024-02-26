package main

import (
	"fmt"
	"log"
	"os"
)

func ExistsOrCreate(path string) *os.File {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("File doesn't exist, creating...")
		file, _ := os.Create(path)
		fmt.Println("File Created")
		return file
	}

	return nil
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}