package cmd

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thaffenden/vrsn/internal/files"
	"github.com/thaffenden/vrsn/internal/flags"
	"github.com/thaffenden/vrsn/internal/git"
	"github.com/thaffenden/vrsn/internal/logger"
	"github.com/thaffenden/vrsn/internal/prompt"
	"github.com/thaffenden/vrsn/internal/version"
)

// NewCmdBump creates the bump command.
func NewCmdBump() *cobra.Command {
	shortDescription := "Increment the current semantic version with a valid patch, major or minor bump."

	cmd := &cobra.Command{
		Args: cobra.OnlyValidArgs,
		RunE: func(ccmd *cobra.Command, args []string) error {
			// TODO: support color option.
			log := logger.NewBasic(false, flags.Verbose)
			curDir, err := os.Getwd()
			if err != nil {
				return err
			}

			log.Debugf("bump command args: %s", args)

			versionFileFinder := files.VersionFileFinder{
				FileFlag:  flags.VersionFile,
				Logger:    log,
				SearchDir: curDir,
			}

			versionFile, err := versionFileFinder.Find()
			if err != nil {
				return err
			}

			currentVersion, err := files.GetVersionFromFile(curDir, versionFile)
			if err != nil {
				return err
			}

			var newVersion string
			if len(args) > 0 {
				options, err := version.GetBumpOptions(currentVersion)
				if err != nil {
					return err
				}

				newVersion, err = options.SelectedIncrement(args[0])
				if err != nil {
					return err
				}
			} else {
				newVersion, err = prompt.SelectBumpType(currentVersion)
				if err != nil {
					return err
				}
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
		Long: fmt.Sprintf(`%s

Pass the increment type directly as an argument to the command, e.g.:

  vrsn bump patch

Or use the interactive prompt to select the increment you want.
The semantic version in the version file will be updated in place.`, shortDescription),
		Short:         shortDescription,
		SilenceErrors: true,
		SilenceUsage:  true,
		Use:           "bump",
		ValidArgs:     []string{"patch", "major", "minor"},
	}

	cmd.Flags().BoolVar(&flags.Commit, "commit", false, "Commit the updated version file after bumping.")
	cmd.Flags().StringVar(&flags.CommitMsg, "commit-msg", "bump version", "Customise the commit message used when committing the version bump.")
	return cmd
}
