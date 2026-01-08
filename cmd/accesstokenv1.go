package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/accesstokenv1"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var accesstokenv1Debug bool

//go:embed descriptions/accesstokenv1.json
var accesstokenv1Descriptions []byte

// accesstokenv1Cmd represents the accesstokenv1 command
var accesstokenv1Cmd = &cobra.Command{
	Use:   "accesstokenv1",
	Short: "Manage Equinix accesstokenv1 resources",
	Long: `Commands for managing Equinix accesstokenv1 resources.

The accesstokenv1 commands are dynamically generated based on the accesstokenv1 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			accesstokenv1.SetDebug(accesstokenv1Debug)
			return accesstokenv1.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(accesstokenv1Cmd)

	// Add common debug flag that will be inherited by all subcommands for accesstokenv1
	accesstokenv1Cmd.PersistentFlags().BoolVar(&accesstokenv1Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for accesstokenv1
	accesstokenv1Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(accesstokenv1Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load accesstokenv1 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := accesstokenv1.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register accesstokenv1 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(accesstokenv1Cmd, client, "accesstokenv1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register accesstokenv1 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// accesstokenv1Cmd.Aliases = []string{"accesstoken"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := accesstokenv1Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
