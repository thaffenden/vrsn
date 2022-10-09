// Package version holds logic for validating an interacting with a single version.
package version

import (
	"strconv"
	"strings"

	"github.com/thaffenden/vrsn/internal/semantic"
	"github.com/thaffenden/vrsn/internal/sentinel"
)

// Validate checks the input string is a valid semantic version and
// parses it into a Version struct.
func Validate(version string) (semantic.Version, error) {
	if !strings.Contains(version, ".") {
		return semantic.Version{}, ErrNoVersionParts
	}

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return semantic.Version{}, ErrNumVersionParts
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return semantic.Version{}, sentinel.WithMessage(err, ErrConvertingToInt, "major version")
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return semantic.Version{}, sentinel.WithMessage(err, ErrConvertingToInt, "minor version")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return semantic.Version{}, sentinel.WithMessage(err, ErrConvertingToInt, "patch version")
	}

	return semantic.Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
