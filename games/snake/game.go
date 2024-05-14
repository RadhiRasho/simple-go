package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
	FoodPos SnakePart
	Score int
	GameOver bool
}

func drawParts(s tcell.Screen, parts []SnakePart, foodPos SnakePart, snakeStyle, foodStyle tcell.Style) {
	s.SetContent(foodPos.x, foodPos.y, '\u25CF', nil, foodStyle)
	for _, part := range parts {
		s.SetContent(part.x, part.y, ' ', nil, snakeStyle)
	}
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, text string) {
	row := y1
	col := x1

	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)

	for _, r := range text {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func checkCollision(parts []SnakePart, otherPart SnakePart) bool {
	for _, part := range parts {
		if part.x == otherPart.x && part.y == otherPart.y {
			return true
		}
	}
	return false
}

func (g *Game) updateFoodPosition(width, height int) {
	g.FoodPos.x = rand.Intn(width)
	g.FoodPos.y = rand.Intn(height)

	if g.FoodPos.y == 1 && g.FoodPos.x < 10 {
		g.updateFoodPosition(width, height)
	}
}

func (g *Game) Run() {
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defaultStyle)
	width, height := g.Screen.Size()
	g.snakeBody.ResetPos(width, height)
	g.updateFoodPosition(width, height)
	g.GameOver = false
	g.Score = 0
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		longerSnake := false

		g.Screen.Clear()

		if checkCollision(g.snakeBody.Parts[len(g.snakeBody.Parts)-1:], g.FoodPos) {
			g.updateFoodPosition(width, height)
			longerSnake = true
			g.Score++
		}
		if checkCollision(g.snakeBody.Parts[:len(g.snakeBody.Parts)-1], g.snakeBody.Parts[len(g.snakeBody.Parts)-1]) {
			break
		}

		g.snakeBody.Update(width, height, longerSnake)
		drawParts(g.Screen, g.snakeBody.Parts, g.FoodPos, snakeStyle, defaultStyle)
		drawText(g.Screen, 1, 1, 8+len(strconv.Itoa(g.Score)), 1, "Score: " + strconv.Itoa(g.Score))
		time.Sleep(60 * time.Millisecond)
		g.Screen.Show()
	}

	g.GameOver = true
	drawText(g.Screen, width/2-20, height/2, width/2+20, height/2, "Game Over, Score: " + strconv.Itoa(g.Score) + " Play Again? y/n")
	g.Screen.Show()
}
