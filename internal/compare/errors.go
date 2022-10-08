package compare

// Error is the error type.
type Error uint

const (
	// ErrNowNotSemVer is the error when both was or now are not supplied.
	ErrNowNotSemVer Error = iota
	// ErrWasNotSemVer is the error when both was or now are not supplied.
	ErrWasNotSemVer
	// ErrNoVersionParts is the error when the version string does not contain any
	// '.' to split into version parts.
	ErrNoVersionParts
	// ErrComparingVersions is the error if the values fall through all of the
	// expected comparison checks.
	ErrComparingVersions
	// ErrComparingVersions is the error is the semantic version does not contain
	// three parts separated by '.'.
	ErrNumVersionParts
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrComparingVersions:
		return "error comparing versions"

	case ErrNumVersionParts:
		return "invalid number of version parts"

	case ErrNowNotSemVer:
		return "now value did not appear to be a semantic version"

	case ErrWasNotSemVer:
		return "was value did not appear to be a semantic version"

	case ErrNoVersionParts:
		return "version string does not contain any . splitting version segments"

	default:
		return "unknown error"
	}
}
