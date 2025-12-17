package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/register"
	"github.com/equinix/cli/internal/stsv1alpha"
	"github.com/spf13/cobra"
)

// TODO: Put service subcommands in separate packages to
// avoid variable naming collisions
var stsv1alphaDebug bool

//go:embed descriptions/stsv1alpha.json
var stsv1alphaDescriptions []byte

// stsv1alphaCmd represents the stsv1alpha command
var stsv1alphaCmd = &cobra.Command{
	Use:   "stsv1alpha",
	Short: "Manage Equinix stsv1alpha resources",
	Long: `Commands for managing Equinix stsv1alpha resources.

The stsv1alpha commands are dynamically generated based on the stsv1alpha API client,
providing access to all available API services.`,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Inject the client factory into the command context so that
		// it can be retrieved at execution time
		executionClientFactory := func() (register.APIClientInterface, error) {
			stsv1alpha.SetDebug(stsv1alphaDebug)
			return stsv1alpha.NewClient()
		}
		cmd.SetContext(register.ContextWithClientFactory(cmd.Context(), executionClientFactory))
	},
}

func init() {
	rootCmd.AddCommand(stsv1alphaCmd)

	// Add common debug flag that will be inherited by all subcommands for stsv1alpha
	stsv1alphaCmd.PersistentFlags().BoolVar(&stsv1alphaDebug, "debug", false, "Enable debug logging for HTTP requests")

	// Add common format flag that will be inherited by all subcommands for stsv1alpha
	stsv1alphaCmd.PersistentFlags().StringP("format", "f", "json", "Format to use for output (json or yaml)")

	// Load SDK descriptions for better command documentation
	if err := register.LoadDescriptions(stsv1alphaDescriptions); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Could not load stsv1alpha descriptions: %v\n", err)
	}

	// Register commands at init time for help/discovery
	// We use a discovery client that doesn't require credentials for structure introspection
	// Actual API calls will validate credentials at runtime
	client, err := stsv1alpha.NewClientForDiscovery()
	if err != nil {
		// If we can't even create a discovery client, log a warning but continue
		// Commands won't be available but the CLI won't crash
		fmt.Fprintf(os.Stderr, "Warning: Could not register stsv1alpha commands: %v\n", err)
		return
	}

	// Register all service commands dynamically
	err = register.ServiceCommands(stsv1alphaCmd, client, "stsv1alpha")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to register stsv1alpha commands: %v\n", err)
	}

	// Uncomment and customize aliases as needed for convenience
	// Example: Add shorter aliases for commonly used commands
	// stsv1alphaCmd.Aliases = []string{"stsv1alpha"}
	//
	// Or add aliases to specific subcommands after registration:
	// if connectionsCmd, _, err := stsv1alphaCmd.Find([]string{"connections"}); err == nil {
	//     if createCmd, _, err := connectionsCmd.Find([]string{"create-connection"}); err == nil {
	//         createCmd.Aliases = []string{"create"}
	//     }
	// }
}
