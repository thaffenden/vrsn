// Package git contains logic for git interactions.
package git

import (
	"bytes"
	"os/exec"

	"github.com/pkg/errors"
)

// CurrentBranch gets the name of the current branch.
func CurrentBranch(dir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = dir

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		return "", errors.Wrapf(err, "error trying to get current git branch name: %s", stdErr.String())
	}

	return stdOut.String(), nil
}
