package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Word struct {
	Word         string   `json:"word"`
	Translations []string `json:"translations"`
}

func UnmarshalWord(data []byte) (Word, error) {
	var r Word
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Word) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

var colorReset, colorRed, colorGreen string = "\033[0m", "\033[31m", "\033[32m"

func main() {
	dir, err := os.ReadDir("./words")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	correct := 0

	fmt.Println("How many would you like to try out? (default: 10)")

	scanner.Scan()

	numWords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Print("Defaulting to 10\n\n")
		numWords = 10
	}

	for i := 0; i < numWords; i++ {
		word := dir[rand.Intn(len(dir))]

		wordData, err := os.ReadFile("./words/" + word.Name())

		if err != nil {
			log.Fatal(err)
		}

		jsonWord, err := UnmarshalWord(wordData)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nThis word means: " + strings.Join(jsonWord.Translations, ", "))

		scanner.Scan()

		input := scanner.Text()

		if strings.EqualFold(strings.TrimSpace(input), jsonWord.Word) {
			fmt.Println(string(colorGreen), "✔ Correct", string(colorReset))
			correct++
		} else {
			fmt.Println(string(colorRed), "❌ Incorrect", string(colorReset))
			fmt.Println(string(colorGreen), "Correct Answer: "+jsonWord.Word, string(colorReset))
		}
	}

	fmt.Println(string(colorGreen), "\nYou Got: " + strconv.Itoa(correct) + " Correct " + "Out of " + strconv.Itoa(numWords), string(colorReset))

}
