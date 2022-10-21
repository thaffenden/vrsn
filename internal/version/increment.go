package version

// GetBumpOptions returns the possible valid version bump options from the
// input string.
func GetBumpOptions(inputVersion string) (map[string]string, error) {
	parsed, err := Parse(inputVersion)
	if err != nil {
		return map[string]string{}, err
	}

	major := parsed
	major.MajorBump()

	minor := parsed
	minor.MinorBump()

	patch := parsed
	patch.PatchBump()

	return map[string]string{
		"major": major.ToString(),
		"minor": minor.ToString(),
		"patch": patch.ToString(),
	}, nil
}
