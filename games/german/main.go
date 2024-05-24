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

	flag.String("challenge", "", "Challenge mode with over 1000 most common words")

	flag.Parse()

	main := strings.ToLower(flag.Arg(0))

	if main == "help" {
		fmt.Println("Usage: german [challenge] | [help] | [no arguments]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var challenge string
	if main == "challenge" {
		challenge = "y"
	}

	scanner := bufio.NewScanner(os.Stdin)

	if challenge == "" {
		fmt.Println("Would you like to take on the Challenge Mode (Y/N)? (Default: N)")
		scanner.Scan()

		challenge = scanner.Text()
	}

	fmt.Println("How many would you like to try out? (default: 10)")

	scanner.Scan()

	numWords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Print("Defaulting to 10\n\n")
		numWords = 10
	}

	correct := 0

	var file []byte

	if strings.ToLower(challenge) == "y" {
		file, err = os.ReadFile("./Top1000.json")
	} else {
		file, err = os.ReadFile("./KnownWords.json")
	}

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	words, err := UnmarshalWords(file)

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	PlayQuiz(words, scanner, numWords, &correct)

	fmt.Println(string(colorGreen), "\nYou Got: "+strconv.Itoa(correct)+" Correct "+"Out of "+strconv.Itoa(numWords), string(colorReset))
}
