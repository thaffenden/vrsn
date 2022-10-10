package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thaffenden/vrsn/internal/files"
	"github.com/thaffenden/vrsn/internal/flags"
	"github.com/thaffenden/vrsn/internal/versions"
)

// NewCmdCheck creates the check command.
func NewCmdCheck() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			// check for git dir
			// if not exists and no was error
			// if exists, and branch is not equal to main
			// 			get files from current branch as now
			// 			get files from main branch as was
			// check for expected version files in directory.
			curDir, err := os.Getwd()
			if err != nil {
				return err
			}

			versionFiles, err := files.GetVersionFilesInDirectory(curDir)
			if err != nil {
				return err
			}

			if len(versionFiles) > 1 {
				return fmt.Errorf("looks like you have several version files: %s", versionFiles)
			}

			if len(versionFiles) == 1 {
				flags.Now, err = files.GetVersionFromFile(curDir, versionFiles[0])
				if err != nil {
					return err
				}
			}

			if err := flags.Validate(flags.Was, flags.Now); err != nil {
				return err
			}

			fmt.Printf("was: %s\nnow: %s\n", flags.Was, flags.Now)

			err = versions.Compare(flags.Was, flags.Now)
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
