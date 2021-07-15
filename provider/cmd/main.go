package main

import (
	"log"

	"github.com/dsinecos/contract_testing/provider"
)

func main() {
	err := provider.InitProvider()
	if err != nil {
		log.Fatalf("Error initializing provider %+v", err)
	}
}
