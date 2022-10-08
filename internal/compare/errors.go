package compare

// Error is the error type.
type Error uint

const (
	// ErrComparingVersions is the error if the values fall through all of the
	// expected comparison checks.
	ErrComparingVersions Error = iota
	// ErrConvertingToInt is the error thrown when a version part cannot be
	// converted to a string.
	ErrConvertingToInt
	// ErrNoVersionParts is the error when the version string does not contain any
	// '.' to split into version parts.
	ErrNoVersionParts
	// ErrNumVersionParts is the error if the semantic version does not contain
	// three parts separated by '.'.
	ErrNumVersionParts
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrComparingVersions:
		return "error comparing versions"

	case ErrConvertingToInt:
		return "error converting version part to int"

	case ErrNumVersionParts:
		return "invalid number of version parts"

	case ErrNoVersionParts:
		return "version string does not contain any . splitting version segments"

	default:
		return "unknown error"
	}
}
