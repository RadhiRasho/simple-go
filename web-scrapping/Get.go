package main

import (
	"global/utils"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func GetRequest() {
	// Make HTTP GET request
	res, err := http.Get("https://www.devdungeon.com/")

	utils.FatalError(err)

	defer res.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, res.Body)

	utils.FatalError(err)

	log.Println("Number of bytes copied to STDOUT: ", n)
}

func GetRequestWithTimeout() {
	// Create HTTP Client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make Request
	res, err := client.Get("https://www.devdungeon.com/")

	utils.FatalError(err)

	defer res.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, res.Body)

	utils.FatalError(err)

	log.Println("Number of bytes copied to STDOUT: ", n)
}


func SetHeaders() {
	// Create HTTP Client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create and Modify HTTP request before sending
	req, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)

	utils.FatalError(err)

	// For web scrapping it's better to have a very discriptive
	// User-Agent to make sure the recieve knows what's going on
	// and can choose to limit your access if it causes issues
	req.Header.Set("User-Agent", "Not Firefox")

	// Make Request
	res, err := client.Do(req)

	utils.FatalError(err)

	// Copy data from the response to standard output
	_, err = io.Copy(os.Stdout, res.Body)

	utils.FatalError(err)
}

func DownloadURL() {
	// Make Request
	res, err := http.Get("https://www.devdungeon.com/archive")

	utils.FatalError(err)

	defer res.Body.Close()

	// Create output file
	outFile, err := os.Create("files/DownloadURL.html")

	utils.FatalError(err)

	defer outFile.Close()

	// Copy Data from HTTP Response to Standard Output
	_, err = io.Copy(outFile, res.Body)

	utils.FatalError(err)
}

func SubstringMatching() {

}