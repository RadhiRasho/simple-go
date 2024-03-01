package main

import (
	"fmt"
	"global/utils"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found
func ProcessElement(index int, element *goquery.Selection, attr string) {
	// See if the href attributes exist on the element
	val, exists := element.Attr(attr)

	if exists {
		fmt.Println(val)
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
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		ProcessElement(index, element, "href")
	})
}

func FindImageInPage() {
	// Make HTTP GET Request
	res, err := http.Get("https://www.devdungeon.com")

	utils.FatalError(err)

	defer res.Body.Close()

	// Create a goquery document from the HTTP Response
	document, err := goquery.NewDocumentFromReader(res.Body)

	utils.FatalError(err)

	document.Find("img").Each(func(index int, element *goquery.Selection) {
		ProcessElement(index, element, "src")
	})
}