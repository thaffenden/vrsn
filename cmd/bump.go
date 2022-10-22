package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thaffenden/vrsn/internal/files"
	"github.com/thaffenden/vrsn/internal/flags"
	"github.com/thaffenden/vrsn/internal/git"
	"github.com/thaffenden/vrsn/internal/logger"
	"github.com/thaffenden/vrsn/internal/prompt"
)

// NewCmdBump creates the bump command.
func NewCmdBump() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			// TODO: support color option.
			log := logger.NewBasic(false, flags.Verbose)
			curDir, err := os.Getwd()
			if err != nil {
				return err
			}

			versionFiles, err := files.GetVersionFilesInDirectory(curDir)
			if err != nil {
				return err
			}

			if len(versionFiles) > 1 {
				return errors.Errorf("looks like you have several version files: %s", versionFiles)
			}

			if len(versionFiles) == 0 {
				return errors.New("no version file found in directory")
			}

			versionFile := versionFiles[0]

			currentVersion, err := files.GetVersionFromFile(curDir, versionFile)
			if err != nil {
				return err
			}

			// TODO: support passing bump type through flag
			newVersion, err := prompt.SelectBumpType(currentVersion)
			if err != nil {
				return err
			}

			if err := files.WriteVersionToFile(curDir, versionFile, newVersion); err != nil {
				return err
			}

			log.Infof("version bumped from %s to %s", currentVersion, newVersion)

			if flags.Commit {
				addOutput, err := git.Add(curDir, versionFile)
				if err != nil {
					return errors.Wrapf(err, "git add output: %s", addOutput)
				}

				commitOutput, err := git.Commit(curDir, versionFile, flags.CommitMsg)
				if err != nil {
					return errors.Wrapf(err, "git add output: %s", commitOutput)
				}

				log.Infof("version file committed")
			}

			return nil
		},
		Short:         "increment semantic version",
		SilenceErrors: true,
		SilenceUsage:  true,
		Use:           "bump",
	}

	cmd.Flags().BoolVar(&flags.Commit, "commit", false, "use this flag to commit the version file after bumping")
	cmd.Flags().StringVar(&flags.CommitMsg, "commit-msg", "bump version", "commit message provided when committing file")
	return cmd
}
