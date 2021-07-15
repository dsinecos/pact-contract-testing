package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Greeting struct {
	Language string `json:"language"`
	Message  string `json:"message"`
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	greeting := Greeting{
		Language: "EN",
		Message:  "Hello",
	}

	serializedGreeting, err := json.Marshal(greeting)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}

	fmt.Fprint(w, string(serializedGreeting))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", greetingHandler)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatalf("Error: %+v", err)
	}
}
