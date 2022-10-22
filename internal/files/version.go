package files

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
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
		"CMakeLists.txt": {
			reader:  getVersionFromCMakeLists,
			updater: updateVersionInCMakeLists,
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

type versionFileMatcher struct {
	lineMatcher   func(string) bool
	versionRegex  string
	notFoundError error
}

func getVersionMatcher(inputFile string) (versionFileMatcher, error) {
	matcher, exists := VersionFileMatchers()[inputFile]
	if !exists {
		return versionFileMatcher{}, fmt.Errorf("%s is not a supported version file type", inputFile)
	}

	return matcher, nil
}

// VersionFileMatchers contains the utilies to extract and update the version
// from the version file.
func VersionFileMatchers() map[string]versionFileMatcher {
	return map[string]versionFileMatcher{
		"Cargo.toml": {
			lineMatcher: func(line string) bool {
				return strings.Contains(line, "version =")
			},
			versionRegex:  `(.*)(version = "){1}(?P<semver>\d+.\d+.\d+)(".*)`,
			notFoundError: ErrGettingVersionFromTOML,
		},
		"CMakeLists.txt": {
			lineMatcher: func(line string) bool {
				return strings.Contains(line, "project(")
			},
			versionRegex:  `(project\(.*)(VERSION ){1}(?P<semver>\d+.\d+.\d+)(.*\))`,
			notFoundError: ErrGettingVersionFromCMakeLists,
		},
	}
}

func (v versionFileMatcher) GetVersion(scanner *bufio.Scanner) (string, error) {
	for scanner.Scan() {
		lineText := scanner.Text()

		if v.lineMatcher(lineText) {
			re := regexp.MustCompile(v.versionRegex)
			result := make(map[string]string)

			match := re.FindStringSubmatch(lineText)
			for i, name := range re.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = match[i]
				}
			}

			fmt.Println(result)

			semver, exists := result["semver"]

			fmt.Printf("SEMVER: %s ||| %v", semver, exists)
			if !exists {
				return "", v.notFoundError
			}

			return semver, nil
		}
	}

	return "", v.notFoundError
}

func (v versionFileMatcher) UpdateVersionInPlace(scanner *bufio.Scanner, newVersion string) ([]string, error) {
	foundVersion := false
	allLines := []string{}

	for scanner.Scan() {
		lineText := scanner.Text()

		if v.lineMatcher(lineText) {
			re := regexp.MustCompile(v.versionRegex)
			newVersionLine := re.ReplaceAllString(lineText, fmt.Sprintf(`${1}${2}%s${4}`, newVersion))
			allLines = append(allLines, newVersionLine)
			foundVersion = true
			continue
		}

		allLines = append(allLines, lineText)
	}

	if !foundVersion {
		return []string{}, v.notFoundError
	}

	return allLines, nil
}
