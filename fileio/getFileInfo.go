package main

import (
	"encoding/json"
	"fmt"
	"global/utils"
	"os"
)


func GetFileInfo() {
	// Stat returns file info. It will return
	// an error if there is no file
	file, err := os.Stat("text-files/truncation.txt")

	utils.FatalError(err)

	fmt.Println("FileName: ", file.Name())
	fmt.Println("Size in bytes: ", file.Size())
	fmt.Println("Permissions: ", file.Mode())
	fmt.Println("Last Modified: ", file.ModTime())
	fmt.Println("Is Directory: ", file.IsDir())
	fmt.Printf("System interface type: %T\n", file.Sys())
	data, _ := json.MarshalIndent(file.Sys(), " ", "    ")
	fmt.Printf("System info: %+v\n\n", string(data))
}