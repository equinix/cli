package cmd

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

// The --config flag has to load the file it points to, not only remember its
// path.
func TestE2E_ConfigFlagLoadsCredentials(t *testing.T) {
	api := newFakeEquinixAPI(t)

	t.Setenv("HOME", t.TempDir())
	t.Setenv("EQUINIX_CLIENT_ID", "")
	t.Setenv("EQUINIX_CLIENT_SECRET", "")
	t.Cleanup(func() {
		cfgFile = ""
		viper.Reset()
	})

	cfg := filepath.Join(t.TempDir(), "custom.yaml")
	if err := os.WriteFile(cfg, []byte("equinix_client_id: test-id\nequinix_client_secret: test-secret\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	rootCmd.SetArgs([]string{
		"--config", cfg,
		"fabricv4", "ip-blocks", "get-ip-block",
		"--uuid", "11111111-2222-3333-4444-555555555555",
	})
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("command failed: %v", err)
	}
	if !api.seen(http.MethodGet, "/fabric/v4/ipBlocks/11111111-2222-3333-4444-555555555555") {
		t.Errorf("expected a GET for the IP block, got %v", api.requests)
	}
}
