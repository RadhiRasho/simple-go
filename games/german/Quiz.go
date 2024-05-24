package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
)

func PlayQuiz(words Words, scanner *bufio.Scanner, numWords int, correct *int) {
	usedWords := make(map[string]bool)

	for i := 0; i < numWords; i++ {
		var word Word
		for {
			word = words[rand.Intn(len(words))]
			if !usedWords[word.Word] {
				break
			}
		}

		usedWords[word.Word] = true

		fmt.Println(string(colorCyan),"\nDefinition: ", string(colorReset), strings.Join(word.Definition, ", "))

		posStrings := make([]string, len(word.Pos))
		for i, pos := range word.Pos {
			posStrings[i] = string(pos)
		}

		fmt.Print(string(colorCyan), "Part of Speech: ", string(colorReset), strings.Join(posStrings, ", "), string(colorReset), "\n")

		scanner.Scan()

		input := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(input), word.Word) {
			fmt.Print(string(colorGreen), " ✔ Correct", string(colorReset), "\n")
			*correct++
		} else {
			fmt.Print(string(colorRed), "❌ Incorrect", string(colorReset), "\n")
			fmt.Print(string(colorGreen), "Correct Answer: ", word.Word, string(colorReset), "\n\n")
			if word.Description != nil {
				fmt.Print(string(colorYellow), "Additional Information: ", string(colorReset), *word.Description, string(colorReset), "\n")
			}
		}
	}
}
