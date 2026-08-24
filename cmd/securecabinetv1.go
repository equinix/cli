package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/register"
	"github.com/equinix/cli/internal/securecabinetv1"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var securecabinetv1Debug bool

//go:embed descriptions/securecabinetv1.json
var securecabinetv1Descriptions []byte

// securecabinetv1Cmd represents the securecabinetv1 command
var securecabinetv1Cmd = &cobra.Command{
	Use:   "securecabinetv1",
	Short: "Manage Equinix securecauinetv1 resources",
	Long: `Commands for managing Equinix securecauinetv1 resources.

The securecabinetv1 commands are dynamically generated based on the securecauinetv1 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			securecabinetv1.SetDebug(securecabinetv1Debug)
			return securecabinetv1.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(securecabinetv1Cmd)

	// Add common debug flag that will be inherited by all subcommands for securecabinetv1
	securecabinetv1Cmd.PersistentFlags().BoolVar(&securecabinetv1Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for securecabinetv1
	securecabinetv1Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(securecabinetv1Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load securecabinetv1 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := securecabinetv1.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register securecabinetv1 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(securecabinetv1Cmd, client, "securecabinetv1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register securecabinetv1 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// securecabinetv1Cmd.Aliases = []string{"securecabinet"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := securecabinetv1Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
