// An implementation of Conway's Game of Life.
// See: https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life
// This is based on the example that is present on the https://go.dev website's main page.
// The example is present at: https://go.dev/play/p/8Yx7J9v6ZvL
// this is a copy that was writtern down slowly in order to understand each part.

package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// Fields represent a two-dimentional field oc cells.
type Field struct {
	cells [][]bool // Two Dimentional Array of booleans
	width, height int // With & h
}


// NewField returns an empty field of the specified width and height.
func NewField(width, height int) *Field {
	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return &Field{cells: cells, width: width, height: height}
}

// SetFieldCell : Sets the state of the specified cell to the given value
func (f *Field) SetFieldCell(x, y int, b bool) {
	f.cells[y][x] = b;
}

// Alive: Checks the cell specified by X & Y Cords to see if it's alive or not
// Reports whether the specified cell is alive
// If the X & Y Cords are outside the field boundaries thery are warped
// toroidally. For instance, an x value of -1 is treated as w-1 (w minus 1).
func (f *Field) Alive(x, y int) bool {
	x += f.width
	x %= f.width
	y += f.height
	y %= f.height
	return f.cells[y][x]
}

// NextCell : return the state of the specified cell at the next time step.
func (f *Field) NextCell(x, y int) bool {
	// Count the adjacent cells that are alive
	alive := 0;

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}

	//* Return next state according to the game rules:
	//* exactly 3 neighbors: on,
	//* exactly 2 neighbors: maintain current state
	//* otherwise off
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

// Life stores teh state of a round of Conways' Game of Life
type Life struct {
	current, next *Field
	width, height int
}

func (l *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			l.next.SetFieldCell(x, y, l.current.NextCell(x, y))
		}
	}
	// Swap fields a and b.
	l.current = l.next
	l.next = l.current
}

// String returns the game board as a string.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.height; y++ {
		for x := 0; x < l.width; x++ {
			b := byte(' ')
			if l.current.Alive(x, y) {
				b = '*' | 0x08
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(width, height int) *Life {
	newCurrent := NewField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		newCurrent.SetFieldCell(rand.Intn(width), rand.Intn(height), true)
	}

	return &Life{
		current: newCurrent, next: NewField(width, height),
		width: width, height: height,
	}
}

func main() {
	l := NewLife(50, 25)
	for {
		go l.Step()
		fmt.Println("\u001b[31m\x0c", l) // Clear screen and print new field.
		time.Sleep(time.Second / 10)
		fmt.Print("\033[2J")
    	fmt.Print("\033[1;1H")
	}
}