package main

import (
	"log"
	"os"
)

func CreateFile() {
	path := "text-files/creation.txt"

	newFile, err := os.Create(path)

	FatalError(err)

	log.Printf("%+v\n", newFile)
	newFile.Close()
}