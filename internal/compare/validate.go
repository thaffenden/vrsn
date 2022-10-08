package compare

import (
	"strings"
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

	// TODO: convert parts to int and return as struct

	return Version{}, nil
}
