package main

import (
	"fmt"
	"log"

	"github.com/dsinecos/contract_testing/consumer"
)

const (
	host = "localhost"
	port = 9000
)

func main() {
	greeting, err := consumer.FetchGreeting(host, port)
	if err != nil {
		log.Fatalf("Error fetching greeting from consumer: %+v", err)
	}

	fmt.Printf("Greeting received from provider - %+v \n", greeting)
}
