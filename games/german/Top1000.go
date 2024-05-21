package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PlayAdvanced(scanner *bufio.Scanner, numWords int, correct *int) {
	file, err := os.ReadFile("./Top1000.json")

	if err != nil {
		log.Fatal(err)
	}

	Words, err := UnmarshalWordAdvanced(file)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numWords; i++ {
		word := Words[0]

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nDefinition: " + word.Definition + " (Pos: " + string(word.Pos) + ")")

		scanner.Scan()

		input := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(input), word.Word) {
			fmt.Println(string(colorGreen), "✔ Correct", string(colorReset))
			*correct++
		} else {
			fmt.Println(string(colorRed), "❌ Incorrect", string(colorReset))
			fmt.Println(string(colorGreen), "Correct Answer: "+word.Word, string(colorReset))
			if word.Description != nil {
				fmt.Println(string(colorYellow), "Additional Information Word: "+*word.Description, string(colorReset))
			}
		}
	}
}
