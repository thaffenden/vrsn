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

	tmpFile, err := os.CreateTemp(dir, "vrsn-tmp-*")
	if err != nil {
		return err
	}

	defer func() {
		if err := tmpFile.Close(); err != nil {
			log.Fatalf("error closing temp file while bumping version: %s\n", err)
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

func updateVersionInPackageJSON(scanner *bufio.Scanner, newVersion string) ([]string, error) {
	return updateVersionInFile(
		scanner,
		`"version": "`,
		`(.*)("version": *"){1}(\d+.\d+.\d+)(".*)`,
		newVersion,
		ErrGettingVersionFromPackageJSON,
	)
}

func updateVersionInTOML(scanner *bufio.Scanner, newVersion string) ([]string, error) {
	return updateVersionInFile(
		scanner,
		`version =`,
		`(.*)(version = "){1}(\d+.\d+.\d+)(".*)`,
		newVersion,
		ErrGettingVersionFromTOML,
	)
}

func updateVersionInVERSIONFile(scanner *bufio.Scanner, newVersion string) ([]string, error) {
	return []string{newVersion}, nil
}

func updateVersionInFile(scanner *bufio.Scanner, versionLine string, regexMatcher string, newVersion string, errorType error) ([]string, error) {
	foundVersion := false
	allLines := []string{}

	for scanner.Scan() {
		lineText := scanner.Text()

		if strings.Contains(lineText, versionLine) {
			re := regexp.MustCompile(regexMatcher)
			newVersionLine := re.ReplaceAllString(lineText, fmt.Sprintf(`${1}${2}%s${4}`, newVersion))
			allLines = append(allLines, newVersionLine)
			foundVersion = true
			continue
		}

		allLines = append(allLines, lineText)
	}

	if !foundVersion {
		return []string{}, errorType
	}

	return allLines, nil
}
