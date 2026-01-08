package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/fabricv4"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var fabricv4Debug bool

//go:embed descriptions/fabricv4.json
var fabricv4Descriptions []byte

// fabricv4Cmd represents the fabricv4 command
var fabricv4Cmd = &cobra.Command{
	Use:   "fabricv4",
	Short: "Manage Equinix fauricv4 resources",
	Long: `Commands for managing Equinix fauricv4 resources.

The fabricv4 commands are dynamically generated based on the fauricv4 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			fabricv4.SetDebug(fabricv4Debug)
			return fabricv4.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(fabricv4Cmd)

	// Add common debug flag that will be inherited by all subcommands for fabricv4
	fabricv4Cmd.PersistentFlags().BoolVar(&fabricv4Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for fabricv4
	fabricv4Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

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
