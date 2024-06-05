package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specials := "!@#$%^&*()_+"

	var length int
	flag.IntVar(&length, "length", 8, "The length of the password")

	var numProbability float64 = 0.2;
	flag.Float64Var(&numProbability, "numProbability", 0.2, "The probability of a number being included in the password")

	var specProbability float64 = 0.2;
	flag.Float64Var(&specProbability, "specProbability", 0.2, "The probability of a special character being included in the password")

	flag.Parse()

	if length < 8 {
		fmt.Println("Password length must be at least 8 characters")
		os.Exit(1)
	}

	if numProbability < 0 || numProbability > 1 {
		fmt.Println("Number probability must be between 0 and 1")
		os.Exit(1)
	}

	if specProbability < 0 || specProbability > 1 {
		fmt.Println("Special character probability must be between 0 and 1")
		os.Exit(1)
	}

	var password string

	for i := 0; i < length; i++ {
		password += string(letters[rand.Intn(len(letters))])

		// Append a number with N% probability (defaults to 20%)
		if rand.Float64() < numProbability {
			password += string(numbers[rand.Intn(len(numbers))])
		}

		// Append a special character with N% probability (defaults to 20%)
		if rand.Float64() < specProbability {
			password += string(specials[rand.Intn(len(specials))])
		}
	}

	fmt.Println(password)
}