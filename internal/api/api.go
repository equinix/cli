package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

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

// NewStandardClient creates a new Client for Equinix APIs that exist under
// api.equinix.com and use OAuth2 client credentials for authentication
func NewStandardClient() (*Client, error) {
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
	client.HTTPClient.Transport = authTransport

	return client, nil
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
