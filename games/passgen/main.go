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

	flag.Parse()

	if length < 8 {
		fmt.Println("Password length must be at least 8 characters")
		os.Exit(1)
	}

	password := ""

	for i := 0; i < length; i++ {
		password += string(letters[rand.Intn(len(letters))])

		// Append a number with 20% probability
		if rand.Float32() < 0.2 {
			password += string(numbers[rand.Intn(len(numbers))])
		}

		// Append a special character with 20% probability
		if rand.Float32() < 0.2 {
			password += string(specials[rand.Intn(len(specials))])
		}
	}

	fmt.Println(password)
}