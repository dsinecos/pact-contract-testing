package main

import (
	"fmt"
	"log"

	"github.com/dsinecos/contract_testing/consumer"
)

const (
	baseURL = "http://localhost:9000"
	path    = "greeting"
)

func main() {
	greeting, err := consumer.FetchGreeting(baseURL, path)
	if err != nil {
		log.Fatalf("Error fetching greeting from consumer: %+v", err)
	}

	fmt.Printf("Greeting received from provider - %+v \n", greeting)
}
