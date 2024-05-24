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

		fmt.Println("\nDefinition: " + strings.Join(word.Definition, ", "))

		posStrings := make([]string, len(word.Pos))
		for i, pos := range word.Pos {
			posStrings[i] = string(pos)
		}

		fmt.Println(string(colorGreen), "Part of Speech: "+strings.Join(posStrings, ", "), string(colorReset))

		scanner.Scan()

		input := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(input), word.Word) {
			fmt.Println(string(colorGreen), "  ✔ Correct", string(colorReset))
			*correct++
		} else {
			fmt.Println(string(colorRed), "❌ Incorrect", string(colorReset))
			fmt.Println(string(colorGreen), "Correct Answer: "+word.Word, string(colorReset))
			if word.Description != nil {
				fmt.Println(string(colorYellow), "Additional Information: "+*word.Description, string(colorReset))
			}
		}
	}
}
