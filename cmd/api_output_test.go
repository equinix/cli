package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newFakeAPIForRawRequests(t *testing.T) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2/v1/token":
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{
				"access_token":  "fake-token",
				"token_timeout": "3600",
				"token_type":    "Bearer",
				"user_name":     "fake-user",
			})
		case "/fabric/v4/routeFilters/rf-1":
			w.WriteHeader(http.StatusNoContent)
		case "/fabric/v4/health":
			w.Header().Set("Content-Type", "text/plain")
			_, _ = w.Write([]byte("OK"))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(srv.Close)

	target, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	orig := http.DefaultTransport
	http.DefaultTransport = redirectTransport{target: target, base: orig}
	t.Cleanup(func() { http.DefaultTransport = orig })
}

// An empty body (204 No Content) used to be printed as "null".
func TestE2E_APIEmptyResponsePrintsNothing(t *testing.T) {
	newFakeAPIForRawRequests(t)

	out, err := runCLI(t, "api", "fabric/v4/routeFilters/rf-1", "-X", "DELETE")
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if out != "" {
		t.Errorf("output = %q, want nothing", out)
	}
}

// A body that is not JSON used to be printed as "null".
func TestE2E_APINonJSONResponseIsPrintedAsIs(t *testing.T) {
	newFakeAPIForRawRequests(t)

	out, err := runCLI(t, "api", "fabric/v4/health", "-X", "GET")
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if out != "OK\n" {
		t.Errorf("output = %q, want %q", out, "OK\n")
	}
}
