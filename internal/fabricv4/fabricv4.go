package fabricv4

import (
	"github.com/equinix/cli/internal/api"
	"github.com/equinix/equinix-sdk-go/services/fabricv4"
)

var debugMode bool

// SetDebug enables or disables debug mode for HTTP requests
func SetDebug(enabled bool) {
	debugMode = enabled
}

// NewClient creates a new Fabric v4 API client using the standard
// authentication configuration from internal/api
func NewClient() (*fabricv4.APIClient, error) {
	// Build client options based on debug mode
	var opts []api.ClientOption
	if debugMode {
		opts = append(opts, api.WithDebug())
	}

	// Use the standard client setup for authentication
	stdClient, err := api.NewStandardClient(opts...)
	if err != nil {
		return nil, err
	}

	// Configure the Fabric v4 client
	configuration := fabricv4.NewConfiguration()
	configuration.HTTPClient = stdClient.HTTPClient
	configuration.AddDefaultHeader("X-SOURCE", "equinix-cli")

	return fabricv4.NewAPIClient(configuration), nil
}

// NewClientForDiscovery creates a Fabric v4 API client for command discovery
// This is used at CLI initialization time to register commands without requiring credentials
func NewClientForDiscovery() (*fabricv4.APIClient, error) {
	// Create a basic client with no authentication for structure discovery only
	configuration := fabricv4.NewConfiguration()
	return fabricv4.NewAPIClient(configuration), nil
}
