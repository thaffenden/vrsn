package files

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// WriteVersionToFile updates the version file with the provided new version
// value.
func WriteVersionToFile(dir string, inputFile string, newVersion string) error {
	versionFunc, err := getVersionFunc(inputFile)
	if err != nil {
		return err
	}

	file, err := os.Open(filepath.Clean(filepath.Join(dir, inputFile)))
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("error closing file: %s", inputFile)
		}
	}()

	scanner := bufio.NewScanner(file)

	if err := versionFunc.writer(scanner, newVersion); err != nil {
		return err
	}

	return nil
}

func writeVersionToPackageJSON(scanner *bufio.Scanner, version string) error {
	return nil
}

func writeVersionToTOML(scanner *bufio.Scanner, version string) error {
	for scanner.Scan() {
		lineText := scanner.Text()

		if strings.Contains(lineText, `version =`) {
			// do write opteration here
			return nil
		}
	}

	return ErrGettingVersionFromTOML
}

func writeVersionToVersionFile(scanner *bufio.Scanner, version string) error {
	return nil
}
