package files

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type versionFileMatcher struct {
	lineMatcher   func(string) bool
	versionRegex  string
	notFoundError error
}

// versionFileMatchers contains the utilies to extract and update the version
// from the version file.
func versionFileMatchers() map[string]versionFileMatcher {
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
		"package.json": {
			lineMatcher: func(line string) bool {
				return strings.Contains(line, `"version": "`)
			},
			versionRegex:  `(.*)("version": *"){1}(?P<semver>\d+.\d+.\d+)(".*)`,
			notFoundError: ErrGettingVersionFromPackageJSON,
		},
		"pyproject.toml": {
			lineMatcher: func(line string) bool {
				return strings.Contains(line, "version =")
			},
			versionRegex:  `(.*)(version = "){1}(?P<semver>\d+.\d+.\d+)(".*)`,
			notFoundError: ErrGettingVersionFromTOML,
		},
		"VERSION": {
			lineMatcher: func(line string) bool {
				// single line file so nothing to match on.
				return true
			},
			versionRegex:  `(.*)(.*)(?P<semver>\d+.\d+.\d+)(.*)`,
			notFoundError: ErrGettingVersionFromVERSION,
		},
	}
}

// getVersionMatcher gets the relevant versionFileMatcher config for the
// provided input file or errors if there is no config for a file with that name.
func getVersionMatcher(inputFile string) (versionFileMatcher, error) {
	matcher, exists := versionFileMatchers()[inputFile]
	if !exists {
		return versionFileMatcher{}, fmt.Errorf("%s is not a supported version file type", inputFile)
	}

	return matcher, nil
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

			semver, exists := result["semver"]
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
