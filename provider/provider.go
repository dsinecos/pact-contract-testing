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

type Error struct {
	HTTPStatusCode int    `json:"status_code"`
	ErrorMessage   string `json:"message"`
}

func internalError(w http.ResponseWriter, r *http.Request) {
	error := Error{
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage:   http.StatusText(http.StatusInternalServerError),
	}

	serializedError, err := json.Marshal(error)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, string(serializedError))
}

func InitProvider() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", greetingHandler)
	mux.HandleFunc("/internalerror", internalError)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
	if err != nil {
		return err
	}
	return nil
}
