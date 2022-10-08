// Package version contains logic for the comparison of the versions.
package version

import (
	"log"
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
func Versions(wasInput string, nowInput string) (ChangeType, error) {
	if wasInput == nowInput {
		return NoIncrement, nil
	}

	was, err := ValidateVersion(wasInput)
	if err != nil {
		return 0, err
	}

	now, err := ValidateVersion(nowInput)
	if err != nil {
		return 0, err
	}

	log.Printf("was: %v\nnow: %v", was, now)

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
