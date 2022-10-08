package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	RunE: func(ccmd *cobra.Command, args []string) error {
		return nil
	},
	Short: "check semantic versions are valid",
	Use:   "check-version",
}

func Execute() error {
	ctx := context.Background()
	return rootCmd.ExecuteContext(ctx)
}
