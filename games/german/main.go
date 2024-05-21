package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Would you like to play Top 1000 (Y/N)? (Default: N)")
	scanner.Scan()

	top1000 := scanner.Text()

	fmt.Println("How many would you like to try out? (default: 10)")

	scanner.Scan()

	numWords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Print("Defaulting to 10\n\n")
		numWords = 10
	}

	correct := 0

	if strings.ToLower(top1000) == "y" {
		PlayAdvanced(scanner, numWords, &correct)
	} else {
		PlayKnownWords(scanner, numWords, &correct)
	}

	fmt.Println(string(colorGreen), "\nYou Got: "+strconv.Itoa(correct)+" Correct "+"Out of "+strconv.Itoa(numWords), string(colorReset))
}
