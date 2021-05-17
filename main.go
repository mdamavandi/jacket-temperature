package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mdamavandi/jacket-temperature/structs"
)

const url = "https://api.tomorrow.io/v4/timelines?location=&fields=temperature&units=imperial&timesteps=1h"

func insertLocale(url string, lat float32, lon float32) string {
	return url[:46] + fmt.Sprintf("%f", lat) + "," + fmt.Sprintf("%f", lon) + url[46:]
}

func main() {
	req, err := http.NewRequest(
		http.MethodGet,
		insertLocale(url, 40.310871, -112.012589),
		nil,
	)
	if err != nil {
		log.Fatalf("error creating HTTP request: %v", err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("apikey", os.Getenv("TOMORROW_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatalf("error reading HTTP response body: %v", err)
	// }

	// log.Println("We got the response:", string(responseBytes))
	var weatherSamples structs.ClimaCellPayload
	if err := json.Unmarshal(responseBytes, &weatherSamples); err != nil {
		log.Fatalf("error deserializing weather data")
	}

	for _, timeline := range weatherSamples.Data.Timelines {
		log.Println("we're here")
		if timeline.Intervals != nil {
			for _, data := range timeline.Intervals {
				log.Printf("The temperature at %s is %f degrees F", data.StartTime, data.Values.Temperature)
			}
		} else {
			log.Printf("No temperature data available between %s and %s", timeline.StartTime, timeline.EndTime)
		}
	}
}
