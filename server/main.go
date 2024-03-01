package main

import (
	"fmt"
	"global/utils"
	"net/http"
)


func main() {
	// API routes

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World from HYM")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := "localhost:3000"
	fmt.Println("Server is running on port ", port)

	// start server on port specified above
	utils.FatalError(http.ListenAndServe(port, nil))
}