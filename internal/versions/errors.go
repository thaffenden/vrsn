package versions

// Error is the error type.
type Error uint

const (
	// ErrInvalidBump is the error when the version has changed but to a value
	// that is not valid sem ver.
	ErrInvalidBump Error = iota + 1
	// ErrVersionNotBumped is the error when the version has not been bumped.
	ErrVersionNotBumped
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrInvalidBump:
		return "invalid version bump"
	case ErrVersionNotBumped:
		return "version has not been bumped"
	default:
		return "unknown error"
	}
}
