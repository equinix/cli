package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/eiav2"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var eiav2Debug bool

//go:embed descriptions/eiav2.json
var eiav2Descriptions []byte

// eiav2Cmd represents the eiav2 command
var eiav2Cmd = &cobra.Command{
	Use:   "eiav2",
	Short: "Manage Equinix eiav2 resources",
	Long: `Commands for managing Equinix eiav2 resources.

The eiav2 commands are dynamically generated based on the eiav2 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			eiav2.SetDebug(eiav2Debug)
			return eiav2.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(eiav2Cmd)

	// Add common debug flag that will be inherited by all subcommands for eiav2
	eiav2Cmd.PersistentFlags().BoolVar(&eiav2Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for eiav2
	eiav2Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(eiav2Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load eiav2 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := eiav2.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register eiav2 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(eiav2Cmd, client, "eiav2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register eiav2 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// eiav2Cmd.Aliases = []string{"eia"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := eiav2Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
