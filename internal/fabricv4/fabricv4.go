package fabricv4

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/equinix/cli/internal/api"
	"github.com/equinix/equinix-sdk-go/services/fabricv4"
)

var debugMode bool

// SetDebug enables or disables debug mode for HTTP requests
func SetDebug(enabled bool) {
	debugMode = enabled
}

// debugTransport wraps an HTTP transport to log requests and responses
type debugTransport struct {
	transport http.RoundTripper
}

func (t *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if debugMode {
		// Log the request
		fmt.Fprintf(os.Stderr, "\n=== HTTP Request ===\n")
		reqDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error dumping request: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", string(reqDump))
		}
	}

	// Execute the request
	resp, err := t.transport.RoundTrip(req)

	if debugMode && resp != nil {
		// Log the response
		fmt.Fprintf(os.Stderr, "\n=== HTTP Response ===\n")
		respDump, dumpErr := httputil.DumpResponse(resp, true)
		if dumpErr != nil {
			fmt.Fprintf(os.Stderr, "Error dumping response: %v\n", dumpErr)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", string(respDump))
		}
		fmt.Fprintf(os.Stderr, "===================\n\n")
	}

	return resp, err
}

// NewClient creates a new Fabric v4 API client using the standard
// authentication configuration from internal/api
func NewClient() (*fabricv4.APIClient, error) {
	// Use the standard client setup for authentication
	stdClient, err := api.NewStandardClient()
	if err != nil {
		return nil, err
	}

	// Wrap the HTTP client with debug transport if needed
	if debugMode {
		transport := stdClient.HTTPClient.Transport
		if transport == nil {
			transport = http.DefaultTransport
		}
		stdClient.HTTPClient.Transport = &debugTransport{transport: transport}
	}

	// Configure the Fabric v4 client
	configuration := fabricv4.NewConfiguration()
	configuration.HTTPClient = stdClient.HTTPClient
	configuration.AddDefaultHeader("X-SOURCE", "equinix-cli")

	if debugMode {
		configuration.Debug = true
		configuration.DefaultHeader["X-Debug"] = "true"
	}

	return fabricv4.NewAPIClient(configuration), nil
}

// NewClientForDiscovery creates a Fabric v4 API client for command discovery
// This is used at CLI initialization time to register commands without requiring credentials
func NewClientForDiscovery() (*fabricv4.APIClient, error) {
	// Create a basic client with no authentication for structure discovery only
	configuration := fabricv4.NewConfiguration()
	return fabricv4.NewAPIClient(configuration), nil
}
