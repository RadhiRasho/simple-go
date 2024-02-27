package main

import (
	"fmt"
	"global/utils"
	"os"
)


func DeleteFile() {
	fmt.Println("File Deletion")

	path := "text-files/deletion.txt"

	ExistsOrCreate(path)

	fmt.Println("Deleting File...")
	// time.Sleep(time.Minute) // Uncomment to see deletion in action after a minute
	err := os.Remove(path)

	utils.FatalError(err)

	fmt.Println("File Deleted Successfully")
}