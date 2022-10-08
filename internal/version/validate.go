package version

import (
	"strconv"
	"strings"

	"github.com/thaffenden/check-version/internal/sentinel"
)

// Version holds the details of the semantic version parts.
type Version struct {
	Major int
	Minor int
	Patch int
}

// ValidateVersion checks the input string is a valid semantic version and
// parses it into a Version struct.
func ValidateVersion(version string) (Version, error) {
	if !strings.Contains(version, ".") {
		return Version{}, ErrNoVersionParts
	}

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return Version{}, ErrNumVersionParts
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return Version{}, sentinel.WithMessage(err, ErrConvertingToInt, "major version")
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return Version{}, sentinel.WithMessage(err, ErrConvertingToInt, "minor version")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return Version{}, sentinel.WithMessage(err, ErrConvertingToInt, "patch version")
	}

	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
