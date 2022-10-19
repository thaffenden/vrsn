// Package cmd contains all of the CLI commands.
package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thaffenden/vrsn/internal/flags"
)

// Version is the CLI version set via linker flags at build time.
var Version string

var rootCmd = &cobra.Command{
	RunE: func(ccmd *cobra.Command, args []string) error {
		err := ccmd.Help()
		if err != nil {
			return err
		}
		return nil
	},
	Short:   "check semantic versions are valid",
	Use:     "vrsn",
	Version: Version,
}

// Execute executes the root command.
func Execute() error {
	ctx := context.Background()
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.AddCommand(NewCmdCheck())
	rootCmd.PersistentFlags().BoolVar(&flags.Verbose, "verbose", false, "get verbose output")
}
