package api

import (
	"context"
	"errors"
	"time"

	"github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/equinix/equinix-sdk-go/services/fabricv4"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/viper"
)

// NewFabricV4Client creates a new Fabric v4 API client with OAuth2 authentication
// and retry capabilities. It uses the client credentials from viper configuration.
func NewFabricV4Client() (*fabricv4.APIClient, error) {
	clientID := viper.GetString("equinix_client_id")
	clientSecret := viper.GetString("equinix_client_secret")

	if clientID == "" || clientSecret == "" {
		return nil, errors.New("equinix_client_id or equinix_client_secret not found in env or config")
	}

	baseURL := "https://api.equinix.com"
	authConfig := equinixoauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      baseURL,
	}
	authTransport := authConfig.New()

	// Set up retryable HTTP client for resilience
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = authTransport
	retryClient.RetryWaitMin = time.Second
	retryClient.RetryWaitMax = time.Second * 60
	standardClient := retryClient.StandardClient()

	// Configure the Fabric v4 client
	configuration := fabricv4.NewConfiguration()
	configuration.HTTPClient = standardClient
	configuration.AddDefaultHeader("X-SOURCE", "equinix-cli")

	return fabricv4.NewAPIClient(configuration), nil
}

// NewFabricV4ClientWithContext creates a new Fabric v4 API client with a context
// This is useful for operations that need cancellation or timeout support
func NewFabricV4ClientWithContext(_ context.Context) (*fabricv4.APIClient, error) {
	// Currently returns the same client as NewFabricV4Client
	// The context can be used in individual API calls
	return NewFabricV4Client()
}

// NewFabricV4ClientForDiscovery creates a Fabric v4 API client for command discovery
// This is used at CLI initialization time to register commands without requiring credentials
// The client uses dummy credentials that won't work for actual API calls
func NewFabricV4ClientForDiscovery() (*fabricv4.APIClient, error) {
	// Create a basic client with no authentication for structure discovery only
	configuration := fabricv4.NewConfiguration()
	return fabricv4.NewAPIClient(configuration), nil
}
