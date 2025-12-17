package stsv1alpha

import (
	"github.com/equinix/cli/internal/api"
	"github.com/equinix/equinix-sdk-go/services/stsv1alpha"
)

var debugMode bool

// SetDebug enables or disables debug mode for HTTP requests
func SetDebug(enabled bool) {
	debugMode = enabled
}

// NewClient creates a new stsv1alpha API client using the standard
// authentication configuration from internal/api
func NewClient() (*stsv1alpha.APIClient, error) {
	// Build client options based on debug mode
	var opts []api.ClientOption
	if debugMode {
		opts = append(opts, api.WithDebug())
	}

	// TEMP: use a Metal client for the metalv1 service
	// until it is removed from public SDKs
	var stdClient *api.Client
	var err error
	if "stsv1alpha" == "metalv1" {
		stdClient, err = api.NewMetalClient(opts...)
	} else {
		// Use the standard client setup for authentication
		stdClient, err = api.NewStandardClient(opts...)
	}

	if err != nil {
		return nil, err
	}

	// Configure the stsv1alpha client
	configuration := stsv1alpha.NewConfiguration()
	configuration.HTTPClient = stdClient.HTTPClient
	configuration.AddDefaultHeader("X-SOURCE", "equinix-cli")

	return stsv1alpha.NewAPIClient(configuration), nil
}

// NewClientForDiscovery creates a stsv1alpha API client for command discovery
// This is used at CLI initialization time to register commands without requiring credentials
func NewClientForDiscovery() (*stsv1alpha.APIClient, error) {
	// Create a basic client with no authentication for structure discovery only
	configuration := stsv1alpha.NewConfiguration()
	return stsv1alpha.NewAPIClient(configuration), nil
}
