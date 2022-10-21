package files

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
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
			log.Fatalf("error closing file: %s\n%s", inputFile, err)
		}
	}()

	scanner := bufio.NewScanner(file)

	newContents, err := versionFunc.updater(scanner, newVersion)
	if err != nil {
		return err
	}

	tmpFile, err := os.CreateTemp(dir, "vrsn*")
	if err != nil {
		return err
	}

	defer func() {
		if err := tmpFile.Close(); err != nil {
			log.Fatal("error closing temp file while bumping version")
		}

		if err := os.Remove(tmpFile.Name()); err != nil {
			log.Fatal("error removing temporary file while bumping version")
		}
	}()

	for _, line := range newContents {
		if _, err := tmpFile.WriteString(fmt.Sprintf("%s\n", line)); err != nil {
			return err
		}
	}

	if err := os.Rename(tmpFile.Name(), file.Name()); err != nil {
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

func updateVersionInTOML(scanner *bufio.Scanner, newVersion string) ([]string, error) {
	foundVersion := false
	allLines := []string{}

	for scanner.Scan() {
		lineText := scanner.Text()

		if strings.Contains(lineText, `version =`) {
			re := regexp.MustCompile(`(.*)(version = "){1}(\d+.\d+.\d+)(".*)`)
			newVersionLine := re.ReplaceAllString(lineText, fmt.Sprintf(`${1}${2}%s${4}`, newVersion))
			allLines = append(allLines, newVersionLine)
			foundVersion = true
			continue
		}

		allLines = append(allLines, lineText)
	}

	if !foundVersion {
		return []string{}, ErrGettingVersionFromTOML
	}

	return allLines, nil
}
