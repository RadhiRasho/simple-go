package main

import (
	"log"
	"os"
)

func CreateFile() {
	newFile, err := os.Create("creation.txt")

	FatalError(err)

	log.Printf("%+v\n", newFile)
	newFile.Close()
}