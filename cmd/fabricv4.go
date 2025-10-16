package cmd

import (
	"fmt"
	"os"

	"github.com/equinix/cli/internal/api"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// fabricv4Cmd represents the fabricv4 command
var fabricv4Cmd = &cobra.Command{
	Use:   "fabricv4",
	Short: "Manage Equinix Fabric v4 resources",
	Long: `Commands for managing Equinix Fabric v4 resources including connections,
cloud routers, networks, ports, service profiles, and more.

The fabricv4 commands are dynamically generated based on the Fabric v4 API client,
providing access to all available API services.`,
	PersistentPreRun: func(_ *cobra.Command, _ []string) {
		// Ensure client is initialized when actually running commands
		// This validates credentials before execution
		_, err := api.NewFabricV4Client()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing Fabric v4 client: %v\n", err)
			fmt.Fprintln(os.Stderr, "\nPlease ensure the following are set:")
			fmt.Fprintln(os.Stderr, "  - EQUINIX_CLIENT_ID environment variable")
			fmt.Fprintln(os.Stderr, "  - EQUINIX_CLIENT_SECRET environment variable")
			fmt.Fprintln(os.Stderr, "\nOr in your config file (~/.config/equinix/equinix.yaml):")
			fmt.Fprintln(os.Stderr, "  equinix_client_id: your-client-id")
			fmt.Fprintln(os.Stderr, "  equinix_client_secret: your-client-secret")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(fabricv4Cmd)

	// Register commands at init time for help/discovery
	// We use a mock client structure just to discover the command structure
	// Actual API calls will validate credentials at runtime
	client, err := api.NewFabricV4ClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register fabricv4 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(fabricv4Cmd, client, "fabricv4")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register fabricv4 commands: %v\n", err)
	}
}

// GetFabricV4Services returns information about all available Fabric v4 services
// This can be useful for documentation or introspection
func GetFabricV4Services() ([]register.ServiceInfo, error) {
	client, err := api.NewFabricV4Client()
	if err != nil {
		return nil, err
	}

	return register.GetServiceList(client)
}
