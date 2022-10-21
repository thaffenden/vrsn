package files

import (
	"bufio"
	"fmt"
)

type (
	readerFunc  func(*bufio.Scanner) (string, error)
	updaterFunc func(*bufio.Scanner, string) ([]string, error)
	writerFunc  func(*bufio.Scanner, string) error
)

type versionFileHandlers struct {
	reader  readerFunc
	updater updaterFunc
	writer  writerFunc
}

// versionFileMap is a map containing the expected name of the version file
// with the function used to extract the version from that file.
func versionFileMap() map[string]versionFileHandlers {
	return map[string]versionFileHandlers{
		"Cargo.toml": {
			reader:  getVersionFromTOML,
			updater: updateVersionInTOML,
			writer:  writeVersionToTOML,
		},
		"package.json": {
			reader:  getVersionFromPackageJSON,
			updater: updateVersionInPackageJSON,
			writer:  writeVersionToPackageJSON,
		},
		"pyproject.toml": {
			reader: getVersionFromTOML,
			writer: writeVersionToTOML,
		},
		"VERSION": {
			reader: getVersionFromVersionFile,
			writer: writeVersionToVersionFile,
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
