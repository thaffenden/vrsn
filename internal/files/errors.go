package files

// Error is the error type.
type Error uint

const (
	// ErrNoVersionFilesInDir is the error when no version files are found.
	ErrNoVersionFilesInDir Error = iota
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrNoVersionFilesInDir:
		return "no version files found in directory"

	default:
		return "unknown error"
	}
}
