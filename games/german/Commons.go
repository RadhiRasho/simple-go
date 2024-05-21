package main

import "encoding/json"

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

var colorReset, colorRed, colorGreen, colorYellow string = "\033[0m", "\033[31m", "\033[32m", "\033[33m"

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
	Word        string  `json:"Word"`
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
