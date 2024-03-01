package main

import (
	"fmt"
	"global/utils"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func ProcessElement(index int, element *goquery.Selection) {
	// See if the href attributes exist on the element
	href, exists := element.Attr("href")

	if exists {
		fmt.Println(href)
	}
}

func GoQueryFindLinksOnPage() {
	// Make HTTP GET Request
	res, err := http.Get("https://www.devdungeon.com")

	utils.FatalError(err)

	defer res.Body.Close()

	// Create a GoQuery document from the HTTP Response
	document, err := goquery.NewDocumentFromReader(res.Body)

	utils.FatalError(err)

	// Find all links and process them with the function
	// defined earlier (ProcessElement())
	document.Find("a").Each(ProcessElement)
}