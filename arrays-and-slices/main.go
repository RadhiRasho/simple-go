package main

import "fmt"

func Arrays() {
	// Initialize a 2D Array
	var grid [10][4]int

	fmt.Println(grid)

	// Another way to initialize a 2D Array
	gridB := [6][4]int{}

	fmt.Println(gridB)
}

func Slices() {
	numRows := 10

	// Initialize a ten length slice of empty slices
	grid := make([][]int, numRows)

	// Verify it is a slice of ten empty slices
	fmt.Println(grid)

	// Initialize those 10 empty slices
	for i := 0; i < numRows; i++ {
		grid[i] = make([]int, 4)
	}

	// grid is a 2d slice of ints with dimensions 10x4
	fmt.Println(grid)
}

func main() {
	Arrays()

	println("\n")

	Slices()
}