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

const (
	baseURL   = "http://%s:%d"
	path      = "greeting"
	errorPath = "internalerror"
)

type Greeting struct {
	Language string `json:"language" pact:"example=EN"`
	Message  string `json:"message" pact:"example=Hello"`
}

func FetchGreeting(host string, port int) (*Greeting, error) {

	httpClient := http.DefaultClient

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s/%s", fmt.Sprintf(baseURL, host, port), path), nil)
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
		return nil, fmt.Errorf(fmt.Sprintf("%s, %s", errorUnmarshallingResponseBody, data), err)
	}

	fmt.Printf("Output %+v", greeting)

	return &greeting, nil
}

type ErrorResponse struct {
	HTTPStatusCode int    `json:"status_code"`
	ErrorMessage   string `json:"message"`
}

type Error struct {
	HTTPStatusCode int
	ErrorMessage   string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error - %s | HTTPStatusCode %d", e.ErrorMessage, e.HTTPStatusCode)
}

func FetchInternalError(host string, port int) (*Greeting, error) {

	httpClient := http.DefaultClient

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, fmt.Sprintf("%s/%s", fmt.Sprintf(baseURL, host, port), errorPath), nil)
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

	var errResponse ErrorResponse
	err = json.Unmarshal(data, &errResponse)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("%s, %s", errorUnmarshallingResponseBody, data), err)
	}

	fmt.Printf("Output %+v", errResponse)

	err = &Error{
		HTTPStatusCode: errResponse.HTTPStatusCode,
		ErrorMessage:   errResponse.ErrorMessage,
	}

	return nil, err
}
