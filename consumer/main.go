package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseURL = "http://localhost:9000"
	path    = "greeting"
)

type Greeting struct {
	Language string `json:"language"`
	Message  string `json:"message"`
}

func main() {
	httpClient := http.DefaultClient

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s/%s", baseURL, path), nil)
	if err != nil {
		log.Fatalf("Error creating request %+v", err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request %+v", err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body %+v", err)
	}

	var greeting Greeting
	err = json.Unmarshal(data, &greeting)
	if err != nil {
		log.Fatalf("Error unmarshalling response body %+v", err)
	}

	fmt.Printf("Greeting received from provider - %+v \n", greeting)
}
