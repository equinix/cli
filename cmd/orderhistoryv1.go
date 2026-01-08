package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/orderhistoryv1"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var orderhistoryv1Debug bool

//go:embed descriptions/orderhistoryv1.json
var orderhistoryv1Descriptions []byte

// orderhistoryv1Cmd represents the orderhistoryv1 command
var orderhistoryv1Cmd = &cobra.Command{
	Use:   "orderhistoryv1",
	Short: "Manage Equinix orderhistoryv1 resources",
	Long: `Commands for managing Equinix orderhistoryv1 resources.

The orderhistoryv1 commands are dynamically generated based on the orderhistoryv1 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			orderhistoryv1.SetDebug(orderhistoryv1Debug)
			return orderhistoryv1.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(orderhistoryv1Cmd)

	// Add common debug flag that will be inherited by all subcommands for orderhistoryv1
	orderhistoryv1Cmd.PersistentFlags().BoolVar(&orderhistoryv1Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for orderhistoryv1
	orderhistoryv1Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(orderhistoryv1Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load orderhistoryv1 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := orderhistoryv1.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register orderhistoryv1 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(orderhistoryv1Cmd, client, "orderhistoryv1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register orderhistoryv1 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// orderhistoryv1Cmd.Aliases = []string{"orderhistory"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := orderhistoryv1Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
