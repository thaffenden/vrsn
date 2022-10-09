// Package files handles logic for interacting with files.
package files

import (
	"os"

	"golang.org/x/exp/slices"
)

// GetVersionFilesInDirectory checks the provided directory for supported
// version files and returns a list of ones found.
func GetVersionFilesInDirectory(dir string) ([]string, error) {
	allFiles, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	versionFiles := []string{}
	supported := supportedVersionFiles()

	for _, file := range allFiles {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if slices.Contains(supported, name) {
			versionFiles = append(versionFiles, name)
		}
	}

	if len(versionFiles) == 0 {
		return []string{}, ErrNoVersionFilesInDir
	}

	return versionFiles, nil
}

func supportedVersionFiles() []string {
	return []string{
		"Cargo.toml",
		"package.json",
		"pyproject.toml",
		"VERSION",
	}
}
