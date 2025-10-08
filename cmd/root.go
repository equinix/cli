/*
Copyright © 2025 Equinix

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "equinix",
	Short:             "Equinix CLI",
	Long:              `Command-line interface for Equinix APIs`,
	DisableAutoGenTag: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/equinix/equinix.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// TODO: Might be best to avoid using the global viper instance
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config/equinix")
		viper.SetConfigType("yaml")
		viper.SetConfigName("equinix")

		err = viper.MergeInConfig()
		// If the config file doesn't exist, ignore it.  There are
		// other ways to configure the CLI, such as env vars, so
		// we don't have to force the existence of a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Another error occurred while reading the config file; handle it as a fatal error
			cobra.CheckErr(err)
		}

		// Read Metal auth token from metal CLI config if it
		// wasn't set in env or equinix config
		if viper.GetString("metal_auth_token") == "" {
			metal := viper.New()
			metalPath := home + "/.config/equinix/metal.yaml"
			metal.SetConfigFile(metalPath)
			err = metal.MergeInConfig()

			// We don't care if there's an error reading the file, since
			// this CLI prefers equinix.yaml and related env vars for config
			if err == nil && metal.GetString("token") != "" {
				viper.Set("metal_auth_token", metal.GetString("token"))
			}
		}
	}

	viper.AutomaticEnv() // read in environment variables that match
}
