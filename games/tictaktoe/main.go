package main

import (
	"log"

	"github.com/gdamore/tcell"
)

type Cells struct {
	x,y int
}

type GameBoard struct {
	Board [][]Cells
}

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	defaultStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite);

	screen.SetStyle(defaultStyle)

	screen.Show()
}