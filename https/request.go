package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Activity struct {
	Activity      string `json:"activity"`
	Type          string `json:"type"`
	Participants  float64  `json:"participants"`
	Price         float64  `json:"price"`
	Link          string `json:"link"`
	Key           string `json:"key"`
	Accessibility float64  `json:"accessibility"`
}

func main() {
	activity := new(Activity)

	data, err := Get(&activity)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
}

func Get(target interface{}) (Activity, error) {
    r, err := http.Get("https://www.boredapi.com/api/activity")
    if err != nil {
        log.Fatal(err)
    }
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	data, err := UnmarshalActivity(body)

	if err != nil {
		log.Fatal(err)
	}

	return data, err
}

func UnmarshalActivity(data []byte) (Activity, error) {
	var r Activity
	err := json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r, err
}

func (r *Activity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
