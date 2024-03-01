package main

import (
	"global/utils"
	"log"
	"net/http"
	"net/url"
)

func PostWithRequestData() {
	res, err := http.PostForm("http://example.com/form", url.Values{
		"username": {"MyUsername"},
		"password": {"123"},
	})

	utils.FatalError(err)

	defer res.Body.Close()

	log.Println(res.Header) // Print the response headers

	// To upload a file, use POST instead of PostForm, provide
	// a content type like application/json or application/octet-stream
	// and then provide an io.Reader with the data

	// http.Post("https://example.com/upload", "image/jpeg", &buff)
}