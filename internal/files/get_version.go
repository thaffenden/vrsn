package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetVersionFromFile reads the version file and returns the semantic
// version contained.
func GetVersionFromFile(dir string, inputFile string) (string, error) {
	extractVersionFunc, exists := versionFileMap()[inputFile]
	if !exists {
		return "", fmt.Errorf("%s is not a supported version file type", inputFile)
	}

	file, err := os.Open(filepath.Clean(filepath.Join(dir, inputFile)))
	if err != nil {
		return "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("error closing file: %s", inputFile)
		}
	}()

	scanner := bufio.NewScanner(file)

	version := extractVersionFunc(scanner)
	return version, nil
}

// versionFileMap is a map containing the expected name of the version file
// with the function used to extract the version from that file.
func versionFileMap() map[string]func(*bufio.Scanner) string {
	return map[string]func(*bufio.Scanner) string{
		"Cargo.toml":     getVersionFromCargoTOML,
		"package.json":   getVersionFromPackageJSON,
		"pyproject.toml": getVersionFromPyprojectTOML,
		"VERSION":        getVersionFromVersionFile,
	}
}

func getVersionFromCargoTOML(scanner *bufio.Scanner) string {
	return ""
}

func getVersionFromPackageJSON(scanner *bufio.Scanner) string {
	return ""
}

func getVersionFromPyprojectTOML(scanner *bufio.Scanner) string {
	return ""
}

func getVersionFromVersionFile(scanner *bufio.Scanner) string {
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}
