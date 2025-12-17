package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/smarthandsv1"
	"github.com/equinix/cli/internal/register"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var smarthandsv1Debug bool

//go:embed descriptions/smarthandsv1.json
var smarthandsv1Descriptions []byte

// smarthandsv1Cmd represents the smarthandsv1 command
var smarthandsv1Cmd = &cobra.Command{
	Use:   "smarthandsv1",
	Short: "Manage Equinix smarthandsv1 resources",
	Long: `Commands for managing Equinix smarthandsv1 resources.

The smarthandsv1 commands are dynamically generated based on the smarthandsv1 API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			smarthandsv1.SetDebug(smarthandsv1Debug)
			return smarthandsv1.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(smarthandsv1Cmd)

	// Add common debug flag that will be inherited by all subcommands for smarthandsv1
	smarthandsv1Cmd.PersistentFlags().BoolVar(&smarthandsv1Debug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for smarthandsv1
	smarthandsv1Cmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(smarthandsv1Descriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load smarthandsv1 descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := smarthandsv1.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register smarthandsv1 commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(smarthandsv1Cmd, client, "smarthandsv1")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register smarthandsv1 commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// smarthandsv1Cmd.Aliases = []string{"smarthands"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := smarthandsv1Cmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
