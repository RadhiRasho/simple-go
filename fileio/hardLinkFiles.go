package main

import (
	"global/utils"
	"os"
)


func HardLinkFiles() {
	// Creating a hard link
	// You will have two file names that point to the same contents
	// changing the contents of one will change the other
	// Deleting/Renaming one will not affect the other
	path := "text-files/HardLink.txt"
	path2 := "text-files/HardLink_Other.txt"
	// Simple read only open. We will cover actually reading
	// and writing to files in examples further down the page
	ExistsOrCreate(path)

	err := os.Link(path, path2)

	utils.FatalError(err)
}
