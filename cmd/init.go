package cmd

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/equinix/cli/internal/api"
	equinixoauth2 "github.com/equinix/equinix-sdk-go/extensions/equinixoauth2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
	"golang.org/x/term"
)

const (
	developerAppsURL = "https://portal.equinix.com/developer-settings/apps"
	developerDocsURL = "https://docs.equinix.com/equinix-api/developer-settings"
)

var (
	initEndpoint    string
	initSTSEndpoint string

	// Overridable in tests.
	openBrowser = defaultOpenBrowser
	verifyToken = defaultVerifyToken
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure API credentials",
	Long: `Interactively configure the client ID and secret the CLI uses to
authenticate with the Equinix API, and save them to the config file.

The credentials are checked by requesting an OAuth2 token before they are
saved. Leaving an answer blank, or pressing Ctrl-C, leaves the config file
unchanged.`,
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runInit(cmd.InOrStdin(), cmd.OutOrStdout())
	},
}

func init() {
	initCmd.Flags().StringVar(&initEndpoint, "api-endpoint", "",
		"Equinix API base URL to save to the config file (env EQUINIX_API_ENDPOINT, default "+api.DefaultEndpoint+")")
	initCmd.Flags().StringVar(&initSTSEndpoint, "sts-endpoint", "",
		"Equinix STS base URL to save to the config file (env EQUINIX_STS_ENDPOINT, default "+api.DefaultSTSEndpoint+")")
	rootCmd.AddCommand(initCmd)
}

func runInit(in io.Reader, out io.Writer) error {
	r := bufio.NewReader(in)
	path := configPath()

	endpoint := strings.TrimRight(initEndpoint, "/")
	if endpoint == "" {
		endpoint = api.Endpoint()
	}

	_, _ = fmt.Fprintf(out, "Config file: %s\n", path)
	if !api.IsDefaultEndpoint(endpoint) {
		_, _ = fmt.Fprintf(out, "API endpoint: %s\n", endpoint)
	} else {
		create, err := prompt(r, out, "Create a new API credential? [Y/n]: ")
		if err != nil {
			return err
		}
		if create == "" || strings.EqualFold(create, "y") || strings.EqualFold(create, "yes") {
			_, _ = fmt.Fprintf(out, "\nOpen %s\n", developerAppsURL)
			if err := openBrowser(developerAppsURL); err != nil {
				_, _ = fmt.Fprintln(out, "(couldn't open a browser; open the link above yourself)")
			}
			_, _ = fmt.Fprintf(out, "Follow %s to create an app,\nthen enter its client ID and secret below.\n\n", developerDocsURL)
		}
	}

	clientID, err := prompt(r, out, "Client ID: ")
	if err != nil {
		return err
	}
	if clientID == "" {
		_, _ = fmt.Fprintln(out, "No client ID entered; config file left unchanged.")
		return nil
	}
	clientSecret, err := promptSecret(in, r, out, "Client Secret: ")
	if err != nil {
		return err
	}
	if clientSecret == "" {
		_, _ = fmt.Fprintln(out, "No client secret entered; config file left unchanged.")
		return nil
	}

	_, _ = fmt.Fprint(out, "Checking credentials... ")
	if err := verifyToken(endpoint, clientID, clientSecret); err != nil {
		_, _ = fmt.Fprintln(out, "failed")
		return fmt.Errorf("could not get an OAuth2 token, config file left unchanged: %w", err)
	}
	_, _ = fmt.Fprintln(out, "ok")

	values := map[string]string{
		"equinix_client_id":     clientID,
		"equinix_client_secret": clientSecret,
	}
	if initEndpoint != "" {
		values[api.EndpointKey] = endpoint
	}
	if initSTSEndpoint != "" {
		values[api.STSEndpointKey] = strings.TrimRight(initSTSEndpoint, "/")
	}
	if err := updateConfigFile(path, values); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(out, "Saved credentials to %s\n", path)
	return nil
}

// prompt prints label and returns the trimmed line the user typed. EOF with
// nothing typed counts as a blank answer.
func prompt(r *bufio.Reader, out io.Writer, label string) (string, error) {
	_, _ = fmt.Fprint(out, label)
	line, err := r.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

// promptSecret is prompt without echoing what's typed, when in is a terminal.
func promptSecret(in io.Reader, r *bufio.Reader, out io.Writer, label string) (string, error) {
	f, ok := in.(*os.File)
	if !ok || !term.IsTerminal(int(f.Fd())) {
		return prompt(r, out, label)
	}
	fd := int(f.Fd())
	state, err := term.GetState(fd)
	if err != nil {
		return "", err
	}

	// ReadPassword turns echo off but leaves Ctrl-C as a signal, which would
	// otherwise exit with the terminal still not echoing.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	done := make(chan struct{})
	defer func() {
		signal.Stop(sigs)
		close(done)
	}()
	go func() {
		select {
		case <-sigs:
			_ = term.Restore(fd, state)
			_, _ = fmt.Fprintln(out)
			os.Exit(130)
		case <-done:
		}
	}()

	_, _ = fmt.Fprint(out, label)
	secret, err := term.ReadPassword(fd)
	_, _ = fmt.Fprintln(out)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(secret)), nil
}

// configPath is the file --config points to, or the default config file.
func configPath() string {
	if cfgFile != "" {
		return cfgFile
	}
	if used := viper.ConfigFileUsed(); used != "" {
		return used
	}
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, ".config", "equinix", "equinix.yaml")
}

// updateConfigFile sets values in the YAML file at path, keeping any other
// keys already there. It doesn't go through viper so that values from the
// environment or flags aren't written out.
func updateConfigFile(path string, values map[string]string) error {
	doc := map[string]any{}
	data, err := os.ReadFile(path)
	switch {
	case err == nil:
		if err := yaml.Unmarshal(data, &doc); err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}
		if doc == nil {
			doc = map[string]any{}
		}
	case !errors.Is(err, os.ErrNotExist):
		return err
	}
	for k, v := range values {
		doc[k] = v
	}
	data, err = yaml.Marshal(doc)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return err
	}
	// WriteFile keeps the mode of an existing file; credentials shouldn't be
	// readable by others.
	return os.Chmod(path, 0o600)
}

func defaultVerifyToken(endpoint, clientID, clientSecret string) error {
	cfg := equinixoauth2.Config{ClientID: clientID, ClientSecret: clientSecret, BaseURL: endpoint}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	tok, err := cfg.TokenSource().TokenWithContext(ctx)
	if err != nil {
		return err
	}
	if tok == nil || tok.AccessToken == "" {
		return errors.New("empty access token")
	}
	return nil
}

func defaultOpenBrowser(url string) error {
	var c *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		c = exec.Command("open", url)
	case "windows":
		c = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		c = exec.Command("xdg-open", url)
	}
	return c.Start()
}
