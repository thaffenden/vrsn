package versions

// Error is the error type.
type Error uint

const (
	// ErrComparingVersions is the error if the values fall through all of the
	// expected comparison checks.
	ErrComparingVersions Error = iota + 1
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrComparingVersions:
		return "error comparing versions"

	default:
		return "unknown error"
	}
}
