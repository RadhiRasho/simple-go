package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type WordAdvanced []WordAdvancedElement

func UnmarshalWordAdvanced(data []byte) (WordAdvanced, error) {
	var r WordAdvanced
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WordAdvanced) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WordAdvancedElement struct {
	Word       string  `json:"Word"`
	Pos         Pos     `json:"pos"`
	Definition  string  `json:"definition"`
	Description *string `json:"description,omitempty"`
}

type Pos string

const (
	Adjective   Pos = "Adjective"
	Adverb      Pos = "Adverb"
	Article     Pos = "Article"
	Conjunction Pos = "Conjunction"
	Determiner  Pos = "Determiner"
	Noun        Pos = "Noun"
	Numeral     Pos = "Numeral"
	Particle    Pos = "Particle"
	Preposition Pos = "Preposition"
	Pronoun     Pos = "Pronoun"
	Propernoun  Pos = "Propernoun"
	Verb        Pos = "Verb"
)

func PlayAdvanced() {
	file, err := os.ReadFile("./Top1000.json")

	if err != nil {
		log.Fatal(err)
	}


	Words, err := UnmarshalWordAdvanced(file)

	if err != nil {
		log.Fatal(err)
	}

	for _, i := range Words {
		fmt.Println(i.Word)
		fmt.Println(i.Pos)
		fmt.Println(i.Definition)
		fmt.Println(i.Description)

	}
}