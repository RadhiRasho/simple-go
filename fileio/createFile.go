package main

import (
	"global/utils"
	"log"
	"os"
)

func CreateFile() {
	path := "text-files/creation.txt"

	newFile, err := os.Create(path)

	utils.FatalError(err)

	log.Printf("%+v\n", newFile)
	newFile.Close()
}