// Package files handles logic for interacting with files.
package files

import (
	"io/fs"
	"os"

	"github.com/pkg/errors"
	"github.com/thaffenden/vrsn/internal/logger"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// VersionFileFinder looks for the relevant version file based on the options
// specified.
type VersionFileFinder struct {
	FileFlag  string
	Logger    logger.Basic
	SearchDir string
}

// Find returns the version file based on the config provided.
func (v VersionFileFinder) Find() (string, error) {
	if v.FileFlag != "" {
		v.Logger.Debugf("using --file flag with file %s", v.FileFlag)

		info, err := os.Stat(v.FileFlag)
		// Handle not exists error first for better user output.
		if errors.Is(err, fs.ErrNotExist) {
			return "", errors.Errorf("file %s not found", v.FileFlag)
		}

		if err != nil {
			return "", errors.Wrapf(err, "error checking for file %s", v.FileFlag)
		}

		if info.IsDir() {
			return "", errors.Errorf("file %s is a directory", v.FileFlag)
		}

		return v.FileFlag, nil
	}

	allVersionFiles, err := GetVersionFilesInDirectory(v.SearchDir)
	if err != nil {
		return "", err
	}

	numberOfVersionFiles := len(allVersionFiles)

	if numberOfVersionFiles == 1 {
		return allVersionFiles[0], nil
	}

	if numberOfVersionFiles == 0 {
		return "", ErrNoVersionFilesInDir
	}

	return "", errors.WithMessage(ErrMultipleVersionFiles, "use the --file CLI flag to pick the specific version file to use")
}

// GetVersionFilesInDirectory checks the provided directory for supported
// version files and returns a list of ones found.
func GetVersionFilesInDirectory(dir string) ([]string, error) {
	allFiles, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	versionFiles := []string{}
	supported := maps.Keys(versionFileMatchers())

	for _, file := range allFiles {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if slices.Contains(supported, name) {
			versionFiles = append(versionFiles, name)
		}
	}

	return versionFiles, nil
}
