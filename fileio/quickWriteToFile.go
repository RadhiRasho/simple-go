package main

import (
	"global/utils"
	"os"
)


func QuickWriteToFile() {
	path := "text-files/quickwrite.txt"

	ExistsOrCreate(path)

	err := os.WriteFile(path, []byte("Hi\n"), 0666)

	utils.FatalError(err)
}
