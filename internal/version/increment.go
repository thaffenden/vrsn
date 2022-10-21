package version

import (
	"fmt"
)

// BumpOptions contains details about the bump options.
type BumpOptions struct {
	Major string
	Minor string
	Patch string
}

// GetBumpOptions returns the possible valid version bump options from the
// input string.
func GetBumpOptions(inputVersion string) (BumpOptions, error) {
	parsed, err := Parse(inputVersion)
	if err != nil {
		return BumpOptions{}, err
	}

	major := parsed
	major.MajorBump()

	minor := parsed
	minor.MinorBump()

	patch := parsed
	patch.PatchBump()

	return BumpOptions{
		Patch: patch.ToString(),
		Minor: minor.ToString(),
		Major: major.ToString(),
	}, nil
}

// PromptOptions returns the options formatted for a user prompt.
func (b BumpOptions) PromptOptions() []string {
	return []string{
		b.formattedPatch(),
		b.formattedMinor(),
		b.formattedMajor(),
	}
}

// SelectedIncrement gets ust the version number from the user selected prompt.
func (b BumpOptions) SelectedIncrement(increment string) string {
	switch increment {
	case b.formattedPatch():
		return b.Patch

	case b.formattedMinor():
		return b.Minor

	case b.formattedMajor():
		return b.Major
	}

	return ""
}

func (b BumpOptions) formattedMajor() string {
	return fmt.Sprintf("major (%s)", b.Major)
}

func (b BumpOptions) formattedMinor() string {
	return fmt.Sprintf("minor (%s)", b.Minor)
}

func (b BumpOptions) formattedPatch() string {
	return fmt.Sprintf("patch (%s)", b.Patch)
}
