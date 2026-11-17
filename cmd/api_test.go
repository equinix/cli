package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// The api command sends --data as the request body, so a JSON document must
// reach the API as that document and not as a JSON string holding it.
func TestE2E_APISendsDataAsGiven(t *testing.T) {
	var gotBody, gotContentType string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/oauth2/v1/token":
			_ = json.NewEncoder(w).Encode(map[string]string{
				"access_token":  "fake-token",
				"token_timeout": "3600",
				"token_type":    "Bearer",
				"user_name":     "fake-user",
			})
		case "/fabric/v4/routers/search":
			b, _ := io.ReadAll(r.Body)
			gotBody = string(b)
			gotContentType = r.Header.Get("Content-Type")
			_, _ = w.Write([]byte(`{"data":[]}`))
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

	payload := `{"filter":{"and":[{"property":"/state","operator":"=","values":["PROVISIONED"]}]}}`
	_, err = runCLI(t, "api", "fabric/v4/routers/search", "-X", "POST", "-d", payload)
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	if gotBody != payload {
		t.Errorf("request body = %s, want %s", gotBody, payload)
	}
	if gotContentType != "application/json" {
		t.Errorf("Content-Type = %q, want application/json", gotContentType)
	}
}
