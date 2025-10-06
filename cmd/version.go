package cmd

import (
	"fmt"

	"github.com/equinix/cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information for this CLI",
	Long:  "Print version information for this CLI",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("equinix version %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
