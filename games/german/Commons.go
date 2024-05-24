package main

import "encoding/json"

var colorReset, colorRed, colorGreen, colorYellow, colorCyan string = "\033[0m", "\033[31m", "\033[32m", "\033[33m", "\033[36m"

type Words []Word

func UnmarshalWords(data []byte) (Words, error) {
	var r Words
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Words) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Word struct {
	Word            string   `json:"word"`
	EnglishSentence *string  `json:"EnglishSentence,omitempty"`
	Pos             []Pos    `json:"pos"`
	Definition      []string `json:"definition"`
	Description     *string  `json:"description,omitempty"`
}

type Pos string

const (
	Abbreviation Pos = "abbreviation"
	Accusative   Pos = "accusative"
	Adjective    Pos = "adjective"
	Adverb       Pos = "adverb"
	Article      Pos = "article"
	Conjunction  Pos = "conjunction"
	Determiner   Pos = "determiner"
	Feminine     Pos = "feminine"
	Infinitive   Pos = "infinitive"
	Interjection Pos = "interjection"
	Masculine    Pos = "masculine"
	Neuter       Pos = "neuter"
	Noun         Pos = "noun"
	Numeral      Pos = "numeral"
	Particle     Pos = "particle"
	PastTense    Pos = "past tense"
	Phrase       Pos = "phrase"
	Plural       Pos = "plural"
	Possessive   Pos = "possessive"
	Preposition  Pos = "preposition"
	Primary      Pos = "primary"
	Pronoun      Pos = "pronoun"
	Secondary    Pos = "secondary"
	Tertiary     Pos = "tertiary"
	Verb         Pos = "verb"
)
