package main

import (
	"io"
	"log"
	"net/http"
)


func DownloadFileOverHttp() {
	path := "text-files/DownloadFileOverHttp.html"

	file := ExistsOrCreate(path)

	defer file.Close()


	// HTTP GET request devdungeon.com
	url := "http://www.devdungeon.com/archive"
	res, err := http.Get(url)

	FatalError(err)

	defer res.Body.Close()

	// Write bytes from HTTP response to file.
	// res.Body satisfies the reader interface.
	// newFile satisfies the writer interface.
	// That allows us to use io.Copy which accepts
	// any type that implements reader adn writer interface
	numBytesWritten, err := io.Copy(file, res.Body)

	FatalError(err)

	log.Printf("Downloaded %d byte file. \n", numBytesWritten)
}