package files

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetVersionFromFile reads the version file and returns the semantic
// version contained.
func GetVersionFromFile(dir string, inputFile string) (string, error) {
	versionFunc, err := getVersionFunc(inputFile)
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

	version, err := versionFunc.reader(scanner)
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

	version, err := versionFunc.reader(scanner)
	if err != nil {
		return "", err
	}
	return version, nil
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
