module web-scrapping

go 1.22.0

replace global/utils => ../utils

require global/utils v0.0.0-00010101000000-000000000000

require (
	github.com/PuerkitoBio/goquery v1.9.1
	github.com/andybalholm/cascadia v1.3.2 // indirect
	golang.org/x/net v0.21.0 // indirect
)
