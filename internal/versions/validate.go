package versions

import (
	"strconv"
	"strings"

	"github.com/thaffenden/check-version/internal/sentinel"
)

// SemVer holds the details of the semantic version parts.
type SemVer struct {
	Major int
	Minor int
	Patch int
}

// Validate checks the input string is a valid semantic version and
// parses it into a Version struct.
func Validate(version string) (SemVer, error) {
	if !strings.Contains(version, ".") {
		return SemVer{}, ErrNoVersionParts
	}

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return SemVer{}, ErrNumVersionParts
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return SemVer{}, sentinel.WithMessage(err, ErrConvertingToInt, "major version")
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return SemVer{}, sentinel.WithMessage(err, ErrConvertingToInt, "minor version")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return SemVer{}, sentinel.WithMessage(err, ErrConvertingToInt, "patch version")
	}

	return SemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}
