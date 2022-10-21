// Package prompt contains logic for prompting user interaction.
package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/thaffenden/vrsn/internal/version"
	"golang.org/x/exp/maps"
)

// SelectBumpType prompts the user to select the type of version increment they
// wish to use.
func SelectBumpType(currentVersion string) (string, error) {
	versionOptions, err := version.GetBumpOptions(currentVersion)
	if err != nil {
		return "", err
	}

	answer := struct {
		Selected string `survey:"bump"`
	}{}

	err = survey.Ask([]*survey.Question{
		{
			Name: "bump",
			Prompt: &survey.Select{
				Message: "select version bump type:",
				Options: maps.Values(versionOptions),
			},
		},
	}, &answer)
	if err != nil {
		return "", err
	}

	return answer.Selected, nil
}
