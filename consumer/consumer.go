package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	errorCreatingRequest           = "error creating request %w"
	errorSendingRequest            = "error sending request %w"
	errorReadingResponseBody       = "error reading response body %w"
	errorUnmarshallingResponseBody = "error unmarshalling response body %w"
)

type Greeting struct {
	Language string `json:"language"`
	Message  string `json:"message"`
}

func FetchGreeting(baseURL, path string) (*Greeting, error) {

	httpClient := http.DefaultClient

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s/%s", baseURL, path), nil)
	if err != nil {
		return nil, fmt.Errorf(errorCreatingRequest, err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(errorSendingRequest, err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(errorReadingResponseBody, err)
	}

	var greeting Greeting
	err = json.Unmarshal(data, &greeting)
	if err != nil {
		return nil, fmt.Errorf(errorUnmarshallingResponseBody, err)
	}

	return &greeting, nil
}

// type Error struct {
// 	IsHTTPError    bool
// 	HTTPStatusCode int
// 	ErrorMessage   string

// 	Err error
// }

// func NewError(
// 	e error, isHTTPError bool, statusCode int, message string,
// ) *Error {
// 	return nil
// }

// func (e *Error) Error() string {
// 	return fmt.Sprintf("Error - %s | HTTPStatusCode %d", e.ErrorMessage, e.HTTPStatusCode)
// }

// func (e *Error) Unwrap() error {
// 	return e.Err
// }
