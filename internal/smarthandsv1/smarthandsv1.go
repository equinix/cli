package smarthandsv1

import (
	"github.com/equinix/cli/internal/api"
	"github.com/equinix/equinix-sdk-go/services/smarthandsv1"
)

var debugMode bool

// SetDebug enables or disables debug mode for HTTP requests
func SetDebug(enabled bool) {
	debugMode = enabled
}

// NewClient creates a new smarthandsv1 API client using the standard
// authentication configuration from internal/api
func NewClient() (*smarthandsv1.APIClient, error) {
	// Build client options based on debug mode
	var opts []api.ClientOption
	if debugMode {
		opts = append(opts, api.WithDebug())
	}

	// TEMP: use a Metal client for the metalv1 service
	// until it is removed from public SDKs
	var stdClient *api.Client
	var err error
	if "smarthandsv1" == "metalv1" {
		stdClient, err = api.NewMetalClient(opts...)
	} else {
		// Use the standard client setup for authentication
		stdClient, err = api.NewStandardClient(opts...)
	}

	if err != nil {
		return nil, err
	}

	// Configure the smarthandsv1 client
	configuration := smarthandsv1.NewConfiguration()
	configuration.HTTPClient = stdClient.HTTPClient
	configuration.AddDefaultHeader("X-SOURCE", "equinix-cli")

	return smarthandsv1.NewAPIClient(configuration), nil
}

// NewClientForDiscovery creates a smarthandsv1 API client for command discovery
// This is used at CLI initialization time to register commands without requiring credentials
func NewClientForDiscovery() (*smarthandsv1.APIClient, error) {
	// Create a basic client with no authentication for structure discovery only
	configuration := smarthandsv1.NewConfiguration()
	return smarthandsv1.NewAPIClient(configuration), nil
}
