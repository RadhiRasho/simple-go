package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

type Game struct {
	time float64
}

func (g *Game) Update() error {
	g.time += math.Pi / 90
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fill background with dark blue color
	screen.Fill(color.RGBA{6, 96, 96, 255})

	// Calculate points and draw them
	for i := 0; i < 40000; i++ {
		x := float64(i % 200)
		y := float64(i / 200)
		px, py := g.calculatePoint(x, y)

		// Only draw if point is within screen bounds
		if px >= 0 && px < screenWidth && py >= 0 && py < screenHeight {
			screen.Set(int(px), int(py), color.RGBA{255, 255, 255, 46})
		}
	}
}

func (g *Game) calculatePoint(x, y float64) (float64, float64) {
	k := x/8 - 12.5
	e := y/8 - 12.5
	mag := math.Sqrt(k*k + e*e)
	o := mag / 12 * math.Cos(math.Sin(k/2)*math.Cos(e/2))
	d := 5 * math.Cos(o)

	// Calculate final x and y coordinates
	newX := (x + d*k*(math.Sin(d*2+g.time)+math.Sin(y*o*o)/9)) / 1.5 + 133
	newY := (y/3 - d*40 + 19*math.Cos(d+g.time)) * 1.5 + 300

	return newX, newY
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Animated Pattern")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}