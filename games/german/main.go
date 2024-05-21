package main

import (
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

	correct := 0;

	for i := 0; i < 10; i++ {
		word := dir[rand.Intn(len(dir))]

		wordData, err := os.ReadFile("./words/" + word.Name())

		if err != nil {
			log.Fatal(err)
		}

		jsonWord, err := UnmarshalWord(wordData)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Your word means: " + strings.Join(jsonWord.Translations, ", "))

		var input string

		_, err = fmt.Scanln(&input)

		if err != nil {
			log.Fatal(err)
		}

		if strings.EqualFold(input, jsonWord.Word) {
			fmt.Println(string(colorGreen), "✔ Correct", string(colorReset))
			correct++;
		} else {
			fmt.Println(string(colorRed), "❌ Incorrect", string(colorReset))
			fmt.Println(string(colorGreen), "Correct Answer: " + jsonWord.Word, string(colorReset))
		}
	}

	fmt.Println(string(colorGreen), "\n\nYou Got: " + strconv.Itoa(correct) + " Correct", string(colorReset))

}
