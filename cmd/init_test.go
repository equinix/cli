package cmd

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

type initCall struct {
	endpoint, id, secret string
}

// setupInit points the config at a temp file and stubs out the browser and
// the token check.
func setupInit(t *testing.T, verifyErr error) (cfg string, opened *[]string, verified *[]initCall) {
	t.Helper()
	cfg = filepath.Join(t.TempDir(), "equinix.yaml")
	cfgFile = cfg
	opened = &[]string{}
	verified = &[]initCall{}

	origOpen, origVerify := openBrowser, verifyToken
	openBrowser = func(u string) error { *opened = append(*opened, u); return nil }
	verifyToken = func(e, id, s string) error {
		*verified = append(*verified, initCall{e, id, s})
		return verifyErr
	}
	t.Setenv("EQUINIX_API_ENDPOINT", "")
	t.Cleanup(func() {
		openBrowser, verifyToken = origOpen, origVerify
		cfgFile = ""
		initEndpoint, initSTSEndpoint = "", ""
		viper.Reset()
	})
	return cfg, opened, verified
}

func readYAML(t *testing.T, path string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	m := map[string]any{}
	if err := yaml.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	return m
}

func TestInit_CreateCredential(t *testing.T) {
	cfg, opened, verified := setupInit(t, nil)
	if err := os.WriteFile(cfg, []byte("other_key: keep\nequinix_client_id: old\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	var out bytes.Buffer
	if err := runInit(strings.NewReader("\nmy-id\nmy-secret\n"), &out); err != nil {
		t.Fatalf("init failed: %v\n%s", err, out.String())
	}

	if len(*opened) != 1 || (*opened)[0] != developerAppsURL {
		t.Errorf("expected browser opened to %s, got %v", developerAppsURL, *opened)
	}
	if !strings.Contains(out.String(), developerDocsURL) {
		t.Errorf("expected docs link in output:\n%s", out.String())
	}
	if want := (initCall{"https://api.equinix.com", "my-id", "my-secret"}); len(*verified) != 1 || (*verified)[0] != want {
		t.Errorf("expected token check %v, got %v", want, *verified)
	}

	got := readYAML(t, cfg)
	if got["equinix_client_id"] != "my-id" || got["equinix_client_secret"] != "my-secret" || got["other_key"] != "keep" {
		t.Errorf("unexpected config: %v", got)
	}
	if _, ok := got["equinix_api_endpoint"]; ok {
		t.Errorf("endpoint shouldn't be saved when not given: %v", got)
	}
	if fi, _ := os.Stat(cfg); fi.Mode().Perm() != 0o600 {
		t.Errorf("expected mode 0600, got %v", fi.Mode().Perm())
	}
}

func TestInit_DeclineCreateStillPromptsForCredentials(t *testing.T) {
	cfg, opened, _ := setupInit(t, nil)

	var out bytes.Buffer
	if err := runInit(strings.NewReader("n\nmy-id\nmy-secret\n"), &out); err != nil {
		t.Fatal(err)
	}
	if len(*opened) != 0 {
		t.Errorf("browser shouldn't open, got %v", *opened)
	}
	if got := readYAML(t, cfg); got["equinix_client_id"] != "my-id" {
		t.Errorf("unexpected config: %v", got)
	}
}

func TestInit_BlankLeavesConfigUnchanged(t *testing.T) {
	for name, input := range map[string]string{
		"blank id":     "y\n\n",
		"blank secret": "y\nmy-id\n\n",
		"eof":          "y\nmy-id\n",
	} {
		t.Run(name, func(t *testing.T) {
			cfg, _, verified := setupInit(t, nil)
			const orig = "equinix_client_id: old\n"
			if err := os.WriteFile(cfg, []byte(orig), 0o600); err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			if err := runInit(strings.NewReader(input), &out); err != nil {
				t.Fatal(err)
			}
			if len(*verified) != 0 {
				t.Errorf("token check shouldn't run")
			}
			if data, _ := os.ReadFile(cfg); string(data) != orig {
				t.Errorf("config changed: %q", data)
			}
		})
	}
}

func TestInit_FailedTokenCheckLeavesConfigUnchanged(t *testing.T) {
	cfg, _, _ := setupInit(t, errors.New("401 Unauthorized"))

	var out bytes.Buffer
	err := runInit(strings.NewReader("y\nmy-id\nbad\n"), &out)
	if err == nil || !strings.Contains(err.Error(), "401") {
		t.Fatalf("expected token error, got %v", err)
	}
	if _, err := os.Stat(cfg); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("config file shouldn't be written")
	}
}

func TestInit_CustomEndpointSkipsInstructions(t *testing.T) {
	cfg, opened, verified := setupInit(t, nil)
	initEndpoint = "https://mockapi.example.com/"
	initSTSEndpoint = "https://mocksts.example.com/"

	var out bytes.Buffer
	// No create question is asked, so the first line is the client ID.
	if err := runInit(strings.NewReader("my-id\nmy-secret\n"), &out); err != nil {
		t.Fatal(err)
	}
	if len(*opened) != 0 || strings.Contains(out.String(), developerDocsURL) {
		t.Errorf("instructions shouldn't be shown for a custom endpoint:\n%s", out.String())
	}
	if want := (initCall{"https://mockapi.example.com", "my-id", "my-secret"}); len(*verified) != 1 || (*verified)[0] != want {
		t.Errorf("expected token check %v, got %v", want, *verified)
	}
	got := readYAML(t, cfg)
	if got["equinix_api_endpoint"] != "https://mockapi.example.com" || got["equinix_sts_endpoint"] != "https://mocksts.example.com" {
		t.Errorf("expected endpoints saved, got %v", got)
	}
}

func TestInit_EndpointFromEnv(t *testing.T) {
	_, opened, verified := setupInit(t, nil)
	t.Setenv("EQUINIX_API_ENDPOINT", "https://mockapi.example.com")
	viper.AutomaticEnv()

	var out bytes.Buffer
	if err := runInit(strings.NewReader("my-id\nmy-secret\n"), &out); err != nil {
		t.Fatal(err)
	}
	if len(*opened) != 0 {
		t.Errorf("browser shouldn't open for a custom endpoint")
	}
	if len(*verified) != 1 || (*verified)[0].endpoint != "https://mockapi.example.com" {
		t.Errorf("expected token check against env endpoint, got %v", *verified)
	}
}

// The real token check and the endpoint override, end to end against the
// fake API.
func TestE2E_InitVerifiesTokenAgainstEndpoint(t *testing.T) {
	api := newFakeEquinixAPI(t)
	// Undo the fake's blanket redirect so only the endpoint setting can route
	// requests to it.
	http.DefaultTransport = http.DefaultTransport.(redirectTransport).base
	t.Setenv("HOME", t.TempDir())
	cfg := filepath.Join(t.TempDir(), "equinix.yaml")
	t.Cleanup(func() {
		cfgFile, initEndpoint = "", ""
		viper.Reset()
	})

	rootCmd.SetIn(strings.NewReader("test-id\ntest-secret\n"))
	rootCmd.SetOut(&bytes.Buffer{})
	t.Cleanup(func() { rootCmd.SetIn(nil); rootCmd.SetOut(nil) })
	rootCmd.SetArgs([]string{"--config", cfg, "init", "--api-endpoint", api.URL})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("init failed: %v", err)
	}
	if !api.seen(http.MethodPost, "/oauth2/v1/token") {
		t.Errorf("expected a token request, got %v", api.requests)
	}

	// Commands then use the saved endpoint.
	t.Setenv("EQUINIX_CLIENT_ID", "")
	t.Setenv("EQUINIX_CLIENT_SECRET", "")
	t.Setenv("EQUINIX_API_ENDPOINT", "")
	viper.Reset()
	rootCmd.SetArgs([]string{"--config", cfg, "fabricv4", "ip-blocks", "get-ip-block",
		"--uuid", "11111111-2222-3333-4444-555555555555"})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if !api.seen(http.MethodGet, "/fabric/v4/ipBlocks/11111111-2222-3333-4444-555555555555") {
		t.Errorf("expected a GET for the IP block, got %v", api.requests)
	}
}
