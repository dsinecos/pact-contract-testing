package provider

import (
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/pact-foundation/pact-go/v2/sugar"
	"github.com/stretchr/testify/assert"
)

// var dir, _ = os.Getwd()
// var pactDir = fmt.Sprintf("../consumer/pacts", dir)
const pactDir = "../consumer/pacts"

// Example Provider Pact: How to run me!
// 1. cd <pact-go>/examples/v3
// 2. go test -v -tags provider .
func Test_Provider(t *testing.T) {
	// SetLogLevel("TRACE")
	// CheckVersion()

	// Start provider API in the background
	go InitProvider()
	// err := InitProvider()
	// if err != nil {
	// 	log.Fatalf("Error initializing Provider %+v", err)
	// }

	verifier := HTTPVerifier{}
	// Authorization middleware
	// This is your chance to modify the request before it hits your provider
	// NOTE: this should be used very carefully, as it has the potential to
	// _change_ the contract
	// f := func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		log.Println("[DEBUG] HOOK request filter")
	// 		// r.Header.Add("Authorization", "Bearer 1234-dynamic-value")
	// 		next.ServeHTTP(w, r)
	// 	})
	// }

	// Verify the Provider with local Pact Files
	err := verifier.VerifyProvider(t, VerifyRequest{
		ProviderBaseURL: "http://localhost:9000",
		PactFiles: []string{
			filepath.ToSlash(fmt.Sprintf("%s/GreetingAPIConsumer-GreetinAPI.json", pactDir)),
		},
		// RequestFilter: f,
		// BeforeEach: func() error {
		// 	log.Println("[DEBUG] HOOK before each")
		// 	return nil
		// },
		// AfterEach: func() error {
		// 	log.Println("[DEBUG] HOOK after each")
		// 	return nil
		// },
		// StateHandlers: StateHandlers{
		// 	"User foo exists": func(setup bool, s ProviderStateV3) (ProviderStateV3Response, error) {

		// 		if setup {
		// 			log.Println("[DEBUG] HOOK calling user foo exists state handler", s)
		// 		} else {
		// 			log.Println("[DEBUG] HOOK teardown the 'User foo exists' state")
		// 		}

		// 		// ... do something, such as create "foo" in the database

		// 		// Optionally (if there are generators in the pact) return provider state values to be used in the verification
		// 		return ProviderStateV3Response{"uuid": "1234"}, nil
		// 	},
		// },
	})

	assert.NoError(t, err)
}
