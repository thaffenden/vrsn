package files

import (
	"bufio"
	"fmt"
)

type (
	readerFunc  func(*bufio.Scanner) (string, error)
	updaterFunc func(*bufio.Scanner, string) ([]string, error)
)

type versionFileHandlers struct {
	reader  readerFunc
	updater updaterFunc
}

// versionFileMap is a map containing the expected name of the version file
// with the function used to extract the version from that file.
func versionFileMap() map[string]versionFileHandlers {
	return map[string]versionFileHandlers{
		"Cargo.toml": {
			reader:  getVersionFromTOML,
			updater: updateVersionInTOML,
		},
		"package.json": {
			reader:  getVersionFromPackageJSON,
			updater: updateVersionInPackageJSON,
		},
		"pyproject.toml": {
			reader:  getVersionFromTOML,
			updater: updateVersionInTOML,
		},
		"VERSION": {
			reader:  getVersionFromVersionFile,
			updater: updateVersionInVERSIONFile,
		},
	}
}

// getVersionFunc gets the relevant version function from the map or errors if
// an unsupported version file is passed.
func getVersionFunc(inputFile string) (versionFileHandlers, error) {
	extractVersionFunc, exists := versionFileMap()[inputFile]
	if !exists {
		return versionFileHandlers{}, fmt.Errorf("%s is not a supported version file type", inputFile)
	}

	return extractVersionFunc, nil
}
