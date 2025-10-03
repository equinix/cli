package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsCmd = &cobra.Command{
	Use:   `docs <destination>`,
	Short: "Generate markdown documentation for this CLI",
	Long:  "Generates markdown documentation for this CLI in the specified directory. Each command gets a markdown file.",
	Example: `  # Generate documentation in the ./docs directory:
  equinix docs ./docs`,

	DisableFlagsInUseLine: true,
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		dest := args[0]
		return doc.GenMarkdownTree(cmd.Parent(), dest)
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
