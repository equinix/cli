package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

type recordingTransport struct{ got []string }

func (r *recordingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r.got = append(r.got, req.URL.String())
	return &http.Response{StatusCode: http.StatusOK, Body: http.NoBody, Request: req}, nil
}

func TestEndpointTransport(t *testing.T) {
	t.Cleanup(viper.Reset)

	for _, tc := range []struct {
		name, api, sts string
		want           []string
	}{
		{
			name: "defaults leave requests alone",
			want: []string{
				"https://api.equinix.com/fabric/v4/ports",
				"https://sts.eqix.equinix.com/use/token",
				"https://example.net/other",
			},
		},
		{
			name: "overrides",
			api:  "http://mockapi.example.com/prefix/",
			sts:  "https://mocksts.example.com",
			want: []string{
				"http://mockapi.example.com/prefix/fabric/v4/ports",
				"https://mocksts.example.com/use/token",
				"https://example.net/other",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			viper.Reset()
			viper.Set(EndpointKey, tc.api)
			viper.Set(STSEndpointKey, tc.sts)

			targets, err := endpointTargets()
			if err != nil {
				t.Fatal(err)
			}
			rec := &recordingTransport{}
			tr := &endpointTransport{targets: targets, transport: rec}
			for _, u := range []string{
				"https://api.equinix.com/fabric/v4/ports",
				"https://sts.eqix.equinix.com/use/token",
				"https://example.net/other",
			} {
				req, _ := http.NewRequest(http.MethodGet, u, nil)
				if _, err := tr.RoundTrip(req); err != nil {
					t.Fatal(err)
				}
			}
			for i := range tc.want {
				if rec.got[i] != tc.want[i] {
					t.Errorf("request %d: got %s, want %s", i, rec.got[i], tc.want[i])
				}
			}
		})
	}
}

func TestEndpointTargetsInvalid(t *testing.T) {
	t.Cleanup(viper.Reset)
	viper.Set(STSEndpointKey, "mocksts.example.com")
	if _, err := endpointTargets(); err == nil {
		t.Error("expected an error for an endpoint without a scheme")
	}
}

func TestRequestKeepsErrorBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/fabric/v4/ports/missing":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`[{"errorCode":"EQ-3142102","errorMessage":"Port not found"}]` + "\n"))
		default:
			w.WriteHeader(http.StatusBadGateway)
		}
	}))
	t.Cleanup(srv.Close)
	c := &Client{BaseURL: srv.URL, HTTPClient: srv.Client()}

	_, err := c.Request("fabric/v4/ports/missing", http.MethodGet, "")
	want := `API error: 404 Not Found: [{"errorCode":"EQ-3142102","errorMessage":"Port not found"}]`
	if err == nil || err.Error() != want {
		t.Errorf("error = %v, want %q", err, want)
	}

	_, err = c.Request("fabric/v4/other", http.MethodGet, "")
	if want := "API error: 502 Bad Gateway"; err == nil || err.Error() != want {
		t.Errorf("error = %v, want %q", err, want)
	}
}
