// Package versions contains logic for the comparison of the versions.
package versions

import (
	"github.com/thaffenden/vrsn/internal/version"
)

// Compare compares the provided versions to see if the increase is a valid
// semver increment.
func Compare(wasInput string, nowInput string) error {
	if wasInput == nowInput {
		return ErrVersionNotBumped
	}

	was, err := version.Validate(wasInput)
	if err != nil {
		return err
	}

	now, err := version.Validate(nowInput)
	if err != nil {
		return err
	}

	if IsValidPatch(was, now) {
		return nil
	}

	if IsValidMinor(was, now) {
		return nil
	}

	if IsValidMajor(was, now) {
		return nil
	}

	return ErrInvalidBump
}
