package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thaffenden/vrsn/internal/files"
	"github.com/thaffenden/vrsn/internal/flags"
	"github.com/thaffenden/vrsn/internal/logger"
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

			// prompt user to select bump type
			// TODO: support passing bump type through flag
			// when version selected, get current version
			// increment by specified version

			currentVersion := "0.9.0"
			newVersion := "1.0.0"
			log.Infof("version bumped from %s to %s", currentVersion, newVersion)

			return nil
		},
		Short:         "increment semantic version",
		SilenceErrors: true,
		SilenceUsage:  true,
		Use:           "bump",
	}
	return cmd
}
