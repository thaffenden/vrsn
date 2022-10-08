// Package compare contains logic for the comparison of the versions.
package compare

import (
	"strings"

	"github.com/pkg/errors"
)

// ChangeType is the type used to hold specifics about the version change.
type ChangeType uint

const (
	// NoIncrement is the response when the values are the same.
	NoIncrement ChangeType = iota + 1
	// ValidPatch is a valid patch version bump.
	ValidPatch
	// ValidMinor is a valid minor version bump.
	ValidMinor
	// ValidMajor is a valid major version bump.
	ValidMajor
)

// Versions compares the provided versions to see if the increase is a valid
// semver increment.
func Versions(was string, now string) (ChangeType, error) {
	if was == now {
		return NoIncrement, nil
	}

	if !strings.Contains(was, ".") {
		return 0, errors.WithMessagef(ErrWasNotSemVer, "was: %s", was)
	}

	if !strings.Contains(now, ".") {
		return 0, errors.WithMessagef(ErrNowNotSemVer, "now: %s", now)
	}

	return 0, ErrComparingVersions
}

// Message is the human friendly string to represent the type of change between
// the versions being compared.
func (c ChangeType) Message() string {
	switch c {
	case NoIncrement:
		return "the supplied values are the same"

	case ValidMajor:
		return "valid major version bump"

	case ValidMinor:
		return "valid minor version bump"

	case ValidPatch:
		return "valid patch version bump"

	default:
		return "invalid change type"
	}
}
