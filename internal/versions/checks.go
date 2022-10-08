package versions

import "github.com/thaffenden/check-version/internal/semantic"

// IsValidMajor checks if the version bump is a valid major bump.
func IsValidMajor(was semantic.Version, now semantic.Version) bool {
	return ((now.Major == was.Major+1) &&
		(now.Minor == 0) &&
		(now.Patch == 0))
}

// IsValidMinor checks if the version bump is a valid minor bump.
func IsValidMinor(was semantic.Version, now semantic.Version) bool {
	return ((now.Major == was.Major) &&
		(now.Minor == was.Minor+1) &&
		(now.Patch == 0))
}

// IsValidPatch checks if the version bump is a valid patch bump.
func IsValidPatch(was semantic.Version, now semantic.Version) bool {
	return ((now.Major == was.Major) &&
		(now.Minor == was.Minor) &&
		(now.Patch == was.Patch+1))
}
