package main

import (
	"encoding/json"
	"global/utils"
	"io"
	"net/http"
)

type GeoIP struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

var (
	address string
	err error
	geo GeoIP
	res *http.Response
	body []byte
)


func main() {
	// Provide a domain name or IP address
	// address = "radhi-rasho.dev"
	// address = "2600:3c00::f03c:91ff:fe98:c0f5" // different variations
	// address = "76.76.21.61" // different variations

	// Use ip-api.com to get a JSON response
	res, err = http.Get("http://ip-api.com/json/" + address)

	utils.FatalError(err)

	defer res.Body.Close()

	// res.Body() is a reader type. We have
	// to use io.ReadAll() to read the data
	// in to a byte slice (string)
	body, err = io.ReadAll(res.Body)

	utils.FatalError(err)

	// Unmarshal the JSON byte slice to a GeoIP struct
	err = json.Unmarshal(body, &geo)

	utils.FatalError(err)

	// Everything accessible in struct now
	println("\n==== IP Geolocation Info ====\n")
	println("IP address:\t", geo.Query)
	println("Country Code:\t", geo.CountryCode)
	println("Country Name:\t", geo.Country)
	println("Zip Code:\t", geo.Zip)
	println("Latitude:\t", geo.Lat)
	println("Longitude:\t", geo.Lon)
	println("Timezone:\t", geo.Timezone)
	println("ISP:\t", geo.ISP)
	println("Org:\t", geo.Org)
}
