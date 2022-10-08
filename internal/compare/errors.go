package compare

// Error is the error type.
type Error uint

const (
	// ErrNowNotSemVer is the error when both was or now are not supplied.
	ErrNowNotSemVer Error = iota
	// ErrWasNotSemVer is the error when both was or now are not supplied.
	ErrWasNotSemVer
	// ErrComparingVersions is the error if the values fall through all of the
	// expected comparison checks.
	ErrComparingVersions
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrComparingVersions:
		return "error comparing versions"

	case ErrNowNotSemVer:
		return "now value did not appear to be a semantic version"

	case ErrWasNotSemVer:
		return "was value did not appear to be a semantic version"

	default:
		return "unknown error"
	}
}
