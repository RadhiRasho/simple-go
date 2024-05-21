package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func PlayKnownWords(scanner *bufio.Scanner, numWords int, correct *int) {
	dir, err := os.ReadDir("./words")

	if err != nil {
		log.Fatal(err)
	}

	usedWords := make(map[string]bool)

	for i := 0; i < numWords; i++ {
		var word os.DirEntry
		for {
			word = dir[rand.Intn(len(dir))]
			if !usedWords[word.Name()] {
				break
			}
		}

		usedWords[word.Name()] = true

		wordData, err := os.ReadFile("./words/" + word.Name())

		if err != nil {
			log.Fatal(err)
		}

		jsonWord, err := UnmarshalWord(wordData)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nTranslations: " + strings.Join(jsonWord.Translations, ", "))

		scanner.Scan()

		input := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(input), jsonWord.Word) {
			fmt.Println(string(colorGreen), "  ✔ Correct", string(colorReset))
			*correct++
		} else {
			fmt.Println(string(colorRed), "❌ Incorrect", string(colorReset))
			fmt.Println(string(colorGreen), "Correct Answer: "+jsonWord.Word, string(colorReset))
		}
	}
}
