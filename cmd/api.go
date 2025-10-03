package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/equinix/cli/internal/api"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var method string
var data string
var format string
var portal bool

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api [url-path]",
	Short: "Make a raw API request to the given Equinix API path",
	Long:  `Make a raw authenticated request to the Equinix API using your configured credentials.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]

		var client *api.Client
		var err error

		if portal {
			client, err = api.NewPortalClient()
		} else if isMetalPath(path) {
			client, err = api.NewMetalClient()
		} else {
			client, err = api.NewStandardClient()
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating API client: %v\n", err)
			os.Exit(1)
		}

		resp, err := client.Request(path, method, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "API request failed: %v\n", err)
			os.Exit(1)
		}

		var raw interface{}
		var pretty []byte
		err = json.Unmarshal(resp, &raw)

		// TODO: move formatting elsewhere, possibly replicating
		// output package from metal CLI
		switch format {
		case "json":
			pretty, err = json.MarshalIndent(raw, "", "  ")
		case "yaml":
			pretty, err = yaml.Marshal(raw)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to print response: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(pretty))
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP method to use")
	apiCmd.Flags().StringVarP(&data, "data", "d", "", "Data to send with POST/PUT requests")
	apiCmd.Flags().StringVarP(&format, "format", "f", "json", "Format to use for output (json or yaml)")
	apiCmd.Flags().BoolVar(&portal, "portal", false, "Use Equinix Portal API (cookie auth)")
}

func isMetalPath(path string) bool {
	metalPattern := regexp.MustCompile(`^/?metal(/.*)?$`)
	return metalPattern.MatchString(path)
}
