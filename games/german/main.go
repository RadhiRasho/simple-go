package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(flag.Args()) > 1 {
		fmt.Println("Can't have more than one argument. Exiting...")
		os.Exit(1)
	}

	flag.String("Help", "", "Help")
	flag.String("Top1000", "", "Top 1000")

	flag.Parse()

	main := strings.ToLower(flag.Arg(0))

	if main == "help" {
		fmt.Println("Usage: german [top1000] | [help] | [no arguments]")
		os.Exit(0)
	}

	var top1000 string
	if main == "top1000" {
		top1000 = "y"
	}

	scanner := bufio.NewScanner(os.Stdin)

	if top1000 == "" {
		fmt.Println("Would you like to play Top 1000 (Y/N)? (Default: N)")
		scanner.Scan()

		top1000 = scanner.Text()
	}

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
