package main

import (
	"encoding/json"
	"fmt"
	"os"
)


func SymLinkFiles() {
	// Creating a symlink
	path := "text-files/SymLink.txt"
	sym := "text-files/SymLink_SYM.txt"

	ExistsOrCreate(path)

	err := os.Symlink(path, sym)

	FatalError(err)

	// LStat will return file info, but if it is actually
	// a symlink, it will return info about the SymLink
	// It will not follow the link and give information
	// about the real file
	// Symlinks do not work in Windows (Running this in WSL2 - UBUNTU Dev Contianer)
	fileInfo, err := os.Lstat(sym)

	FatalError(err)

	data, err := json.MarshalIndent(fileInfo, " ", "	")

	FatalError(err)

	fmt.Printf("Link info: %+v", data)

	// Change ownership of a symlink only
	// and not the file it points to
	err = os.Lchown(sym, os.Geteuid(), os.Getgid())
	FatalError(err)
}
