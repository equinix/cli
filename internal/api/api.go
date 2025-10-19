// Package api provides a set of Equinix API clients that can be used
// to make arbitrary requests to Equinix APIs.  This is useful for
// integrating with API endpoints that are not mentioned in the
// Equinix API catalog (https://docs.equinix.com/api-catalog).
// Inspired by https://github.com/displague/equinix-api/ which
// implements the same functionality in bash.
package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	equinixoauth2 "github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/spf13/viper"
)

var standardHeaders = map[string]string{
	"User-Agent": "equinix-cli/version generic-request-client",
}

// Client is a generic Equinix API client that can be used for any Equinix API.
type Client struct {
	BaseURL        string
	DefaultHeaders map[string]string
	HTTPClient     *http.Client
}

// debugTransport wraps an HTTP transport to log requests and responses when debug mode is enabled
type debugTransport struct {
	transport http.RoundTripper
}

func (t *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log the request
	fmt.Fprintf(os.Stderr, "\n==================== HTTP REQUEST ====================\n")
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dumping request: %v\n", err)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", string(reqDump))
	}
	fmt.Fprintf(os.Stderr, "======================================================\n")

	// Execute the request
	resp, err := t.transport.RoundTrip(req)

	if resp != nil {
		// Log the response
		fmt.Fprintf(os.Stderr, "\n==================== HTTP RESPONSE ====================\n")
		respDump, dumpErr := httputil.DumpResponse(resp, true)
		if dumpErr != nil {
			fmt.Fprintf(os.Stderr, "Error dumping response: %v\n", dumpErr)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", string(respDump))
		}
		fmt.Fprintf(os.Stderr, "=======================================================\n\n")
	}

	return resp, err
}

// NewStandardClient creates a new Client for Equinix APIs that exist under
// api.equinix.com and use OAuth2 client credentials for authentication
func NewStandardClient(options ...ClientOption) (*Client, error) {
	client := &Client{
		BaseURL:        "https://api.equinix.com",
		DefaultHeaders: standardHeaders,
		HTTPClient:     &http.Client{},
	}

	clientID := viper.GetString("equinix_client_id")
	clientSecret := viper.GetString("equinix_client_secret")
	if clientID == "" || clientSecret == "" {
		return nil, errors.New("client_id or client_secret not found in env or config")
	}

	authConfig := equinixoauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		BaseURL:      client.BaseURL,
	}
	authTransport := authConfig.New()
	
	// Apply options to potentially wrap the transport
	transport := http.RoundTripper(authTransport)
	for _, opt := range options {
		transport = opt(transport)
	}
	
	client.HTTPClient.Transport = transport

	return client, nil
}

// ClientOption is a function that can modify the HTTP transport
type ClientOption func(http.RoundTripper) http.RoundTripper

// WithDebug returns a ClientOption that enables debug logging of HTTP requests and responses
func WithDebug() ClientOption {
	return func(transport http.RoundTripper) http.RoundTripper {
		return &debugTransport{transport: transport}
	}
}

// NewPortalClient creates a new Client for Equinix APIs that exist under
// portal.equinix.com and rely on Cookies to transmit OAuth2 tokens
func NewPortalClient() (*Client, error) {
	client := &Client{
		BaseURL:        "https://portal.equinix.com/api",
		DefaultHeaders: standardHeaders,
		HTTPClient:     &http.Client{},
	}

	cookie := viper.GetString("equinix_portal_cookie")
	if cookie == "" {
		return nil, errors.New("portal cookie not found in env or config")
	}
	client.DefaultHeaders["Cookie"] = cookie
	client.DefaultHeaders["User-Agent"] = fmt.Sprintf("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:142.0) Gecko/20100101 Firefox/142.0 %s", client.DefaultHeaders["User-Agent"])
	client.DefaultHeaders["Accept"] = "*/*"
	client.DefaultHeaders["Accept-Encoding"] = "*/*"

	return client, nil
}

// NewMetalClient creates a new Client for the Equinix Metal API
func NewMetalClient() (*Client, error) {
	client := &Client{
		BaseURL:        "https://api.equinix.com",
		DefaultHeaders: standardHeaders,
		HTTPClient:     &http.Client{},
	}

	token := viper.GetString("metal_auth_token")
	if token == "" {
		return nil, errors.New("metal_auth_token not found in env or config")
	}
	client.DefaultHeaders["X-Auth-Token"] = token
	return client, nil
}

// Request makes an HTTP request to the specified API path with the given method and data.
//
// TODO: May be useful to support debug logging of requests/responses.  That could
// be done here but probably better to do it with a custom Transport that is shared
// across generic and generated clients for a consistent debug experience.
func (c *Client) Request(apiPath, method string, data string) ([]byte, error) {
	url := c.BaseURL + "/" + apiPath
	var body io.Reader
	if data != "" {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range c.DefaultHeaders {
		req.Header.Set(k, v)
	}
	if data != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}
	return io.ReadAll(resp.Body)
}
