package fabricv4

import (
	"github.com/equinix/cli/internal/api"
	"github.com/equinix/equinix-sdk-go/services/fabricv4"
)

// NewClient creates a new Fabric v4 API client using the standard
// authentication configuration from internal/api
func NewClient() (*fabricv4.APIClient, error) {
	// Use the standard client setup for authentication
	stdClient, err := api.NewStandardClient()
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
