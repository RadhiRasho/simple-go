package main

import (
	"fmt"
	"global/utils"
	"os"
)

func RenameFile() {
	fmt.Println("File Rename")

	originalPath := "text-files/Rename.txt"

	ExistsOrCreate(originalPath)

	newPath := "text-files/Rename2.txt"

	fmt.Println("Renaming", originalPath, "to ", newPath)
	err := os.Rename(originalPath, newPath)

	utils.FatalError(err)

	fmt.Println("Rename complete")
}