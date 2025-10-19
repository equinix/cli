package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/fabricv4"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

//go:embed descriptions/fabricv4.json
var fabricv4Descriptions []byte

// fabricv4Cmd represents the fabricv4 command
var fabricv4Cmd = &cobra.Command{
	Use:   "fabricv4",
	Short: "Manage Equinix Fabric v4 resources",
	Long: `Commands for managing Equinix Fabric v4 resources including connections,
cloud routers, networks, ports, service profiles, and more.

The fabricv4 commands are dynamically generated based on the Fabric v4 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Ensure client is initialized when actually running commands
		// This validates credentials before execution
		debug, _ := cmd.Flags().GetBool("debug")
		fabricv4.SetDebug(debug)

		_, err := fabricv4.NewClient()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing Fabric v4 client: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(fabricv4Cmd)

	// Add debug flag
	fabricv4Cmd.PersistentFlags().Bool("debug", false, "Enable debug mode to show HTTP requests and responses")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(fabricv4Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load fabricv4 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := fabricv4.NewClientForDiscovery()
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

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// fabricv4Cmd.Aliases = []string{"fabric"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := fabricv4Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
