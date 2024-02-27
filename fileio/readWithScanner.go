package main

import (
	"bufio"
	"fmt"
	"global/utils"
	"log"
	"os"
)

func ReadWithScanner() {
	path := "text-files/ReadWithScanner.txt"

	ExistsOrCreate(path)

	file, err := os.Open(path)

	utils.FatalError(err)

	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanWords)

	// Scan for next token.
	success := scanner.Scan()

	if !success {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err != nil {
			log.Println("Scan Completed and Reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// Get data from scan with Bytes() or Text()
	fmt.Println("First word found: ", scanner.Text())
	// Call scanner.Scan() again to find next token
}
