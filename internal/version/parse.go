// Package version holds logic for validating an interacting with a single version.
package version

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thaffenden/vrsn/internal/sentinel"
)

// SemVer holds the details of the semantic version parts.
type SemVer struct {
	Major int
	Minor int
	Patch int
}

// Parse checks the input string is a valid semantic version and
// parses it into a SemVer struct.
func Parse(version string) (SemVer, error) {
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

// MajorBump increments the major version by 1.
func (s *SemVer) MajorBump() {
	s.Major++
	s.Minor = 0
	s.Patch = 0
}

// MinorBump increments the minor version by 1.
func (s *SemVer) MinorBump() {
	s.Minor++
	s.Patch = 0
}

// PatchBump increments the patch version by 1.
func (s *SemVer) PatchBump() {
	s.Patch++
}

// ToString returns the string representation of a SemVer struct.
func (s SemVer) ToString() string {
	return fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
}
