package consumer

import (
	"testing"

	. "github.com/pact-foundation/pact-go/v2/sugar"
	"github.com/stretchr/testify/assert"
)

func Test_FetchGreeting(t *testing.T) {

	mockProvider, err := NewV2Pact(MockHTTPProviderConfig{
		Consumer: "GreetingAPIConsumer",
		Provider: "GreetinAPI",
	})
	assert.NoError(t, err)

	// Arrange: Setup our expected interactions
	mockProvider.
		AddInteraction().
		Given("An endpoint to fetch greeting exists").
		UponReceiving("A request for Greeting").
		WithRequest("GET", S("/greeting")).
		WillRespondWith(200).
		WithBodyMatch(&Greeting{})

	// Act: test our API client behaves correctly
	err = mockProvider.ExecuteTest(t, func(config MockServerConfig) error {
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
