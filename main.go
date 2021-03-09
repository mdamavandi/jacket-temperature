package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "https://data.climacell.co/v4/timelines?location=&fields=temperature&timesteps=1h"

func insertLocale(url string, lat float32, lon float32) string {
	return url[:48] + fmt.Sprintf("%f", lat) + "," + fmt.Sprintf("%f", lon) + url[48:]
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
	req.Header.Add("apikey", os.Getenv("CLIMACELL_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}

	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	log.Println("We got the response:", string(responseBytes))
}
