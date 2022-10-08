// Package cmd contains all of the CLI commands.
package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thaffenden/check-version/internal/flags"
	"github.com/thaffenden/check-version/internal/versions"
)

var rootCmd = &cobra.Command{
	RunE: func(ccmd *cobra.Command, args []string) error {
		if err := flags.Validate(flags.Was, flags.Now); err != nil {
			return err
		}

		_, err := versions.Compare(flags.Was, flags.Now)
		if err != nil {
			return err
		}

		return nil
	},
	Short:            "check semantic versions are valid",
	TraverseChildren: true,
	Use:              "check-version",
}

// Execute executes the root command.
func Execute() error {
	ctx := context.Background()
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.Flags().StringVar(&flags.Was, "was", "", "the previous semantic version (if passing for direct comparison)")
	rootCmd.Flags().StringVar(&flags.Now, "now", "", "the current semantic version (if passing for direct comparison)")
}
