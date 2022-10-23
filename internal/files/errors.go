package files

// Error is the error type.
type Error uint

const (
	// ErrNoVersionFilesInDir is the error when no version files are found.
	ErrNoVersionFilesInDir Error = iota
	// ErrGettingVersionFromCMakeLists is the error when the version can't be
	// found inside a CMakeLists.txt file.
	ErrGettingVersionFromCMakeLists
	// ErrGettingVersionFromPackageJSON is the error when a version key can't be
	// found inside a package.json file.
	ErrGettingVersionFromPackageJSON
	// ErrGettingVersionFromSetupPy is the error when a version key can't be found
	// inside a setup.py file.
	ErrGettingVersionFromSetupPy
	// ErrGettingVersionFromTOML is the error when a version key can't be found
	// inside a toml file.
	ErrGettingVersionFromTOML
	// ErrGettingVersionFromVERSION is the error when the VERSION file is empty.
	ErrGettingVersionFromVERSION
)

// Error returns the error string for the error enum.
func (e Error) Error() string {
	switch e {
	case ErrNoVersionFilesInDir:
		return "no version files found in directory"

	case ErrGettingVersionFromCMakeLists:
		return "error getting version from CMakeLists.txt"

	case ErrGettingVersionFromPackageJSON:
		return "error getting version from package.json"

	case ErrGettingVersionFromSetupPy:
		return "error getting version from setup.py"

	case ErrGettingVersionFromTOML:
		return "error getting version from toml file"

	case ErrGettingVersionFromVERSION:
		return "error getting version from VERSION file"

	default:
		return "unknown error"
	}
}
