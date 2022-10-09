package files

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	version, err := extractVersionFunc(scanner)
	if err != nil {
		return "", err
	}
	return version, nil
}

// versionFileMap is a map containing the expected name of the version file
// with the function used to extract the version from that file.
func versionFileMap() map[string]func(*bufio.Scanner) (string, error) {
	return map[string]func(*bufio.Scanner) (string, error){
		"Cargo.toml":     getVersionFromTOML,
		"package.json":   getVersionFromPackageJSON,
		"pyproject.toml": getVersionFromTOML,
		"VERSION":        getVersionFromVersionFile,
	}
}

func getVersionFromTOML(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		lineText := scanner.Text()

		if strings.Contains(lineText, `version =`) {
			return strings.Split(lineText, `"`)[1], nil
		}
	}
	return "", errors.New("error getting version from Cargo.toml")
}

func getVersionFromPackageJSON(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.Contains(lineText, `"version": "`) {
			return strings.Split(lineText, `"`)[3], nil
		}
	}

	return "", errors.New("error getting version from package.json")
}

func getVersionFromVersionFile(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		// Single line file, so can just return on first line.
		return scanner.Text(), nil
	}

	return "", errors.New("error getting version from VERSION file")
}
