package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/lookupv2"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var lookupv2Debug bool

//go:embed descriptions/lookupv2.json
var lookupv2Descriptions []byte

// lookupv2Cmd represents the lookupv2 command
var lookupv2Cmd = &cobra.Command{
	Use:   "lookupv2",
	Short: "Manage Equinix lookupv2 resources",
	Long: `Commands for managing Equinix lookupv2 resources.

The lookupv2 commands are dynamically generated based on the lookupv2 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			lookupv2.SetDebug(lookupv2Debug)
			return lookupv2.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(lookupv2Cmd)

	// Add common debug flag that will be inherited by all subcommands for lookupv2
	lookupv2Cmd.PersistentFlags().BoolVar(&lookupv2Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for lookupv2
	lookupv2Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(lookupv2Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load lookupv2 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := lookupv2.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register lookupv2 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(lookupv2Cmd, client, "lookupv2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register lookupv2 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// lookupv2Cmd.Aliases = []string{"lookup"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := lookupv2Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
