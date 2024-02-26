package main

import (
	"fmt"
	"os"
)

func RenameFile() {
	fmt.Println("File Rename")

	originalPath := "Rename.txt"

	ExistsOrCreate(originalPath)

	newPath := "Rename2.txt"

	fmt.Println("Renaming", originalPath, "to ", newPath)
	err := os.Rename(originalPath, newPath)

	FatalError(err)

	fmt.Println("Rename complete")
}