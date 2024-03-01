package main

import (
	"fmt"
	"global/utils"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
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
	// Make HTTP GET Request
	res, err := http.Get("Https://www.devdungeon.com")

	utils.FatalError(err)

	defer res.Body.Close()

	// Get the response body as a string
	dataInBytes, err := io.ReadAll(res.Body)

	utils.FatalError(err)

	pageContent := string(dataInBytes)

	// Find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")

	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}

	titleEndIndex := strings.Index(pageContent, "</title>")

	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}

	// (Optional)
	// Copy the substring in to a separate variable so the
	// variable with the full document data can be garbage collected
	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	// Print out the result
	fmt.Printf("Page title: %s\n", pageTitle)
}

func RegExMatching() {
	// Make HTTP GET Request
	res, err := http.Get("https://www.devdungeon.com")

	utils.FatalError(err)

	defer res.Body.Close()

	// Read response data into memory
	body, err := io.ReadAll(res.Body)

	utils.FatalError(err)

	// Create a regular expresison to find comments
	regex := regexp.MustCompile("<!--(.|\n)*?-->")

	comments := regex.FindAllString(string(body), -1)

	if comments == nil {
		fmt.Println("No Matches")
		os.Exit(0)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func ParseURLS() {
	// Parse a complex URL
	complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"

	parsedUrl, err := url.Parse(complexUrl)

	utils.FatalError(err)

	// Print out URL prices
	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Host: ", parsedUrl.Host)
	fmt.Println("Path: ", parsedUrl.Path)
	fmt.Println("Query String: ", parsedUrl.RawQuery)
	fmt.Println("fragment: ", parsedUrl.Fragment)

	// Get the query key/values as a map
	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = "google.com"
	newQueryValues := customURL.Query()
	newQueryValues.Set("ME", "YOU")
	newQueryValues.Set("YOU", "ME")
	customURL.Fragment = "BookmarkLink"
	customURL.RawQuery = newQueryValues.Encode()

	fmt.Println("\nCustom URL: ")
	fmt.Println(customURL.String())
}
