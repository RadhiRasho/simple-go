package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)


func HashingFile() {
	path := "text-files/HashingFile.txt"

	ExistsOrCreate(path)

	// get bytes from file
	data, err := os.ReadFile(path)

	FatalError(err)

	// Hash the file and output result
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))
}

func ChecksumFiles() {
	path := "text-files/ChecksumFiles.txt"

	ExistsOrCreate(path)

	// Open file for reading
	file, err := os.Open(path)

	FatalError(err)

	defer file.Close()

	// Create new hasher, which is a writer interface
	hasher := md5.New()

	_, err = io.Copy(hasher, file)

	FatalError(err)

	// Hash and print. Pass nil since
	// the data is not coming in as a slice argument
	// but is coming through the writer interface
	sum := hasher.Sum(nil)

	fmt.Printf("Md5 checksum: %x\n", sum)
}