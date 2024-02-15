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
	activity := new(Activity);

	data := Get(&activity)

	fmt.Printf("%+v\n", data)
}

func Get(target interface{}) Activity {
    r, err := http.Get("https://www.boredapi.com/api/activity")
    if err != nil {
        log.Fatal(err)
    }
    defer r.Body.Close()

    body, err := io.ReadAll(r.Body);

	if err != nil {
		log.Fatal(err);
	}

	data := UnmarshalActivity(body)

	return data;
}

func UnmarshalActivity(data []byte) Activity {
	var r Activity
	err := json.Unmarshal(data, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func (r *Activity) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
