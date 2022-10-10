package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type versionFunc func(*bufio.Scanner) (string, error)

// GetVersionFromFile reads the version file and returns the semantic
// version contained.
func GetVersionFromFile(dir string, inputFile string) (string, error) {
	extractVersionFunc, err := getVersionFunc(inputFile)
	if err != nil {
		return "", err
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

// GetVersionFromString handles extracting the version from an file that has
// already been read and is passed as a string such as when getting the
// contents of a file from a git branch.
func GetVersionFromString(fileName string, input string) (string, error) {
	versionFunc, err := getVersionFunc(fileName)
	if err != nil {
		return "", err
	}

	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	version, err := versionFunc(scanner)
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

// getVersionFunc gets the relevant version function from the map or errors if
// an unsupported version file is passed.
func getVersionFunc(inputFile string) (versionFunc, error) {
	extractVersionFunc, exists := versionFileMap()[inputFile]
	if !exists {
		return nil, fmt.Errorf("%s is not a supported version file type", inputFile)
	}

	return extractVersionFunc, nil
}

func getVersionFromTOML(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		lineText := scanner.Text()

		if strings.Contains(lineText, `version =`) {
			return strings.Split(lineText, `"`)[1], nil
		}
	}
	return "", ErrGettingVersionFromTOML
}

func getVersionFromPackageJSON(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		lineText := scanner.Text()
		if strings.Contains(lineText, `"version": "`) {
			return strings.Split(lineText, `"`)[3], nil
		}
	}

	return "", ErrGettingVersionFromPackageJSON
}

func getVersionFromVersionFile(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		// Single line file, so can just return on first line.
		return scanner.Text(), nil
	}

	return "", ErrGettingVersionFromVERSION
}
