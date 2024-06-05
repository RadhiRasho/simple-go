package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var red string = "\033[31m"
var green string = "\033[32m"
var yellow string = "\033[33m"
var blue string = "\033[34m"
var purple string = "\033[35m"
var cyan string = "\033[36m"
var white string = "\033[37m"
var reset string = "\033[0m"

var colors = []string{red, green, yellow, blue, purple, cyan, white}

func main() {
	var fileName string

	flag.StringVar(&fileName, "file", "tree", "File to read")

	flag.Parse()

	file, err := os.Open("./" + fileName + ".txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		for _, i := range data {
			fmt.Print(colors[rand.Intn(len(colors))], string(i), reset)
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println()
	}
}
