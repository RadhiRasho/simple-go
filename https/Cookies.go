package main

import (
	"fmt"
	"global/utils"
	"net/http"
)

// Cookie functions for Request
// https://golang.org/pkg/net/http/#Request
// Request.AddCookie()  // Add cookie to request
// Request.Cookie()     // Get specific cookie
// Request.Cookies()    // Get all cookies

// Cookie functions for Response
// https://golang.org/pkg/net/http/#Response
// Response.Cookies()   // Get all cookies

func RequestWithCookies() {
	req, err := http.NewRequest("GET", "https:/www.devdungeon.com", nil)

	utils.FatalError(err)

	// Create a new cookie with the only required feilds
	myCookie := &http.Cookie{
		Name: "Cookie1",
		Value: "Value1",
	}

	// Add the cookie to your request
	req.AddCookie(myCookie)

	// Ask the request to tell us about itself
	// just to confirm the cookie attached properly
	fmt.Println(req.Cookies())
	fmt.Println(req.Header)

	// Do something with the request
	// client := &http.client{}
	// client.Do(req)
}