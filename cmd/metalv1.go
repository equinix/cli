package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/metalv1"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var metalv1Debug bool

//go:embed descriptions/metalv1.json
var metalv1Descriptions []byte

// metalv1Cmd represents the metalv1 command
var metalv1Cmd = &cobra.Command{
	Use:   "metalv1",
	Short: "Manage Equinix metalv1 resources",
	Long: `Commands for managing Equinix metalv1 resources.

The metalv1 commands are dynamically generated based on the metalv1 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			metalv1.SetDebug(metalv1Debug)
			return metalv1.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(metalv1Cmd)

	// Add common debug flag that will be inherited by all subcommands for metalv1
	metalv1Cmd.PersistentFlags().BoolVar(&metalv1Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for metalv1
	metalv1Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(metalv1Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load metalv1 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := metalv1.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register metalv1 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(metalv1Cmd, client, "metalv1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register metalv1 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// metalv1Cmd.Aliases = []string{"metal"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := metalv1Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
