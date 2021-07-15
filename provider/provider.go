package provider

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	PORT = 9000
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
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, string(serializedGreeting))
}

func InitProvider() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", greetingHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
	if err != nil {
		return err
	}
	return nil
}
