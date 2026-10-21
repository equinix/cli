package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"sync"
	"testing"
)

// redirectTransport sends every outgoing request to the test server,
// regardless of the host the SDK built the URL for.
type redirectTransport struct {
	target *url.URL
	base   http.RoundTripper
}

func (rt redirectTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.URL.Scheme = rt.target.Scheme
	req.URL.Host = rt.target.Host
	return rt.base.RoundTrip(req)
}

// fakeEquinixAPI is a minimal stand-in for api.equinix.com.
type fakeEquinixAPI struct {
	*httptest.Server

	mu       sync.Mutex
	requests []string // "METHOD /path" for each request received
}

func (f *fakeEquinixAPI) seen(method, path string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, r := range f.requests {
		if r == method+" "+path {
			return true
		}
	}
	return false
}

func newFakeEquinixAPI(t *testing.T) *fakeEquinixAPI {
	t.Helper()
	f := &fakeEquinixAPI{}
	f.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		f.requests = append(f.requests, r.Method+" "+r.URL.Path)
		f.mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/oauth2/v1/token":
			_ = json.NewEncoder(w).Encode(map[string]string{
				"access_token":  "fake-token",
				"token_timeout": "3600",
				"token_type":    "Bearer",
				"user_name":     "fake-user",
			})
		case r.Method == http.MethodGet && r.URL.Path == "/fabric/v4/ipBlocks/11111111-2222-3333-4444-555555555555":
			if r.Header.Get("Authorization") != "Bearer fake-token" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// The SDK validates required response properties, so the fixture
			// has to include them even though the tests don't care about them.
			_ = json.NewEncoder(w).Encode(map[string]any{
				"uuid":         "11111111-2222-3333-4444-555555555555",
				"href":         "https://api.equinix.com/fabric/v4/ipBlocks/11111111-2222-3333-4444-555555555555",
				"type":         "IPV4_IP_BLOCK",
				"state":        "ACTIVE",
				"ownership":    "EQUINIX",
				"prefixLength": 29,
				"project":      map[string]string{"projectId": "p1"},
				"changeLog":    map[string]string{"createdBy": "fake-user", "createdDateTime": "2025-01-01T00:00:00Z"},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(f.Close)

	target, err := url.Parse(f.URL)
	if err != nil {
		t.Fatal(err)
	}
	orig := http.DefaultTransport
	http.DefaultTransport = redirectTransport{target: target, base: orig}
	t.Cleanup(func() { http.DefaultTransport = orig })

	return f
}

// runCLI executes the real root command with the given args against the fake
// API and returns whatever the command printed to stdout.
func runCLI(t *testing.T, args ...string) (string, error) {
	t.Helper()

	// Isolate from any real user config and supply fake credentials.
	t.Setenv("HOME", t.TempDir())
	t.Setenv("EQUINIX_CLIENT_ID", "test-id")
	t.Setenv("EQUINIX_CLIENT_SECRET", "test-secret")

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	origStdout := os.Stdout
	os.Stdout = w
	var buf bytes.Buffer
	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(&buf, r)
		close(done)
	}()

	rootCmd.SetArgs(args)
	runErr := rootCmd.Execute()

	os.Stdout = origStdout
	_ = w.Close()
	<-done
	return buf.String(), runErr
}

// Regression test: the command name for OAuth2TokenApi ("oauth2-token") can't
// be mapped back to the SDK field naively, which used to panic with
// "reflect: Call using zero Value argument".
func TestE2E_GetOAuth2AccessToken(t *testing.T) {
	api := newFakeEquinixAPI(t)

	out, err := runCLI(t,
		"accesstokenv1", "oauth2-token", "get-oauth2-access-token",
		"--payload-grant_type", "client_credentials",
		"--payload-client_id", "test-id",
		"--payload-client_secret", "test-secret",
	)
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if !api.seen(http.MethodPost, "/oauth2/v1/token") {
		t.Errorf("expected a request to POST /oauth2/v1/token, got %v", api.requests)
	}
	if !strings.Contains(out, "fake-token") {
		t.Errorf("expected output to contain the access token, got:\n%s", out)
	}
}

// Exercises a service whose SDK name contains an initialism (IPBlocksApi) and
// that sends a path parameter and authenticated request.
func TestE2E_FabricGetIPBlock(t *testing.T) {
	api := newFakeEquinixAPI(t)

	out, err := runCLI(t,
		"fabricv4", "ip-blocks", "get-ip-block",
		"--uuid", "11111111-2222-3333-4444-555555555555",
	)
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if !api.seen(http.MethodGet, "/fabric/v4/ipBlocks/11111111-2222-3333-4444-555555555555") {
		t.Errorf("expected a GET for the IP block, got %v", api.requests)
	}
	if !strings.Contains(out, "11111111-2222-3333-4444-555555555555") {
		t.Errorf("expected output to contain the IP block uuid, got:\n%s", out)
	}
}
