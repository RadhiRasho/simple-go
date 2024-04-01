package main

import (
	"fmt"
	"time"
)

type Game struct { temp int }

func CelsiusToFahrenheit(name string, game chan *Game) {
	for {
		game2 := <-game
		game2.temp = int(float64(1.8) * float64(game2.temp))
		fmt.Println(name, game2.temp)
		time.Sleep(100 * time.Millisecond)
		game <- game2
	}
}

func FahrenheitToCelsius(name string, game chan *Game) {
	for {
		game1 := <-game
		game1.temp = (game1.temp * 9/5) + 32
		fmt.Println(name, game1.temp)
		time.Sleep(100 * time.Millisecond)
		game <- game1
	}
}

func main() {
	game := make(chan *Game);

	go CelsiusToFahrenheit("Fahrenheit", game)
	go FahrenheitToCelsius("Celsius", game)

	game <- new(Game)
    time.Sleep(1 * time.Second)
    <-game
}