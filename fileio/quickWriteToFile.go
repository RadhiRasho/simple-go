package main

import "os"


func QuickWriteToFile() {
	path := "quickwrite.txt"

	ExistsOrCreate(path)

	err := os.WriteFile(path, []byte("Hi\n"), 0666)

	FatalError(err)
}
