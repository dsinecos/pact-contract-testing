package consumer

import (
	"fmt"
	"net/http"
	"testing"

	pact "github.com/pact-foundation/pact-go/v2/sugar"
	"github.com/stretchr/testify/assert"
)

func Test_FetchGreeting(t *testing.T) {

	mockProvider, err := pact.NewV3Pact(pact.MockHTTPProviderConfig{
		Consumer: "GreetingAPIConsumer",
		Provider: "GreetinAPI",
	})
	assert.NoError(t, err)

	// Arrange: Setup our expected interactions
	mockProvider.
		AddInteraction().
		Given(pact.ProviderStateV3{
			Name: "An endpoint to fetch greeting exists",
		}).
		UponReceiving("A request for Greeting").
		WithRequest("GET", pact.S("/greeting")).
		WillRespondWith(200).
		WithHeader("Content-Type", pact.S("application/json")).
		// WithJSONBody(pact.Equality("test")) // Returns error "attempted to use matchers from a higher spec version". Need to use pact.NewV3Pact to use these matchers
		WithJSONBody(pact.Map{
			"language": pact.S("EN"),
			"message":  pact.S("Hello"),
		})
		// WithJSONBody(Greeting{})
		// WithJSONBody(pact.Term("admin", "admin|user|guest"))
		// WithBodyMatch(Greeting{})

	// Act: test our API client behaves correctly
	err = mockProvider.ExecuteTest(t, func(config pact.MockServerConfig) error {
		// Initialise the API client and point it at the Pact mock server
		// client := newClient(config.Host, config.Port)

		// Execute the API client
		greeting, err := FetchGreeting(config.Host, config.Port)

		// Assert: check the result
		assert.NoError(t, err)
		assert.Equal(t, &Greeting{Language: "EN", Message: "Hello"}, greeting)

		return err
	})
	assert.NoError(t, err)
}

func Test_FetchInternalError(t *testing.T) {

	mockProvider, err := pact.NewV3Pact(pact.MockHTTPProviderConfig{
		Consumer: "GreetingAPIConsumer",
		Provider: "GreetinAPI",
	})
	assert.NoError(t, err)

	// Arrange: Setup our expected interactions
	mockProvider.
		AddInteraction().
		Given(pact.ProviderStateV3{
			Name: "An endpoint to fetch internal error exists",
		}).
		UponReceiving("A request that triggers internal server error").
		WithRequest("GET", pact.S("/internalerror")).
		WillRespondWith(500).
		WithHeader("Content-Type", pact.S("application/json")).
		WithJSONBody(pact.Map{
			"status_code": pact.Integer(500),
			"message":     pact.S("Internal Server Error"),
		})

	// Act: test our API client behaves correctly
	err = mockProvider.ExecuteTest(t, func(config pact.MockServerConfig) error {
		// Execute the API client
		greeting, err := FetchInternalError(config.Host, config.Port)

		// Assert: check the result
		assert.EqualError(t, err, fmt.Sprintf("Error - %s | HTTPStatusCode %d", http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError))

		assert.Nil(t, greeting)

		return nil
	})

	// assert.Error(t, err)
	assert.NoError(t, err)
}
