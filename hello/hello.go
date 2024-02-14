package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of the logger, including
	//the log entry prefix and a flag to disable printing
	//the time, source file, and line number.
	log.SetPrefix("Hello: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("H")
	if err != nil {
		log.Fatal(err)
	}

	log.SetPrefix("Hellos: ")
	log.SetFlags(0)

	names := []string{
		"R",
		"H",
	}

	messages, err := greetings.Hellos(names)

	// If an err was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)

	for idx, message := range messages {
		fmt.Println(idx)
		fmt.Println(message)
	}
}
