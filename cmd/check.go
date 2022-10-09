package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thaffenden/vrsn/internal/flags"
	"github.com/thaffenden/vrsn/internal/versions"
)

// NewCmdCheck creates the check command.
func NewCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			// check for expected version files in directory.
			if err := flags.Validate(flags.Was, flags.Now); err != nil {
				return err
			}

			fmt.Printf("was: %s\nnow: %s\n", flags.Was, flags.Now)

			err := versions.Compare(flags.Was, flags.Now)
			if err != nil {
				return err
			}

			fmt.Printf("valid version bump\n")

			return nil
		},
		Short:         "check semantic versions are valid",
		SilenceErrors: true,
		SilenceUsage:  true,
		Use:           "check",
	}
	cmd.Flags().StringVar(&flags.Was, "was", "", "the previous semantic version (if passing for direct comparison)")
	cmd.Flags().StringVar(&flags.Now, "now", "", "the current semantic version (if passing for direct comparison)")
	return cmd
}
