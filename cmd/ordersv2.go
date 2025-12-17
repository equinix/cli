package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/ordersv2"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var ordersv2Debug bool

//go:embed descriptions/ordersv2.json
var ordersv2Descriptions []byte

// ordersv2Cmd represents the ordersv2 command
var ordersv2Cmd = &cobra.Command{
	Use:   "ordersv2",
	Short: "Manage Equinix ordersv2 resources",
	Long: `Commands for managing Equinix ordersv2 resources.

The ordersv2 commands are dynamically generated based on the ordersv2 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			ordersv2.SetDebug(ordersv2Debug)
			return ordersv2.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(ordersv2Cmd)

	// Add common debug flag that will be inherited by all subcommands for ordersv2
	ordersv2Cmd.PersistentFlags().BoolVar(&ordersv2Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for ordersv2
	ordersv2Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(ordersv2Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load ordersv2 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := ordersv2.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register ordersv2 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(ordersv2Cmd, client, "ordersv2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register ordersv2 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// ordersv2Cmd.Aliases = []string{"orders"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := ordersv2Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
