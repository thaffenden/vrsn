// Package git contains logic for git interactions.
package git

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// CurrentBranch gets the name of the current branch.
func CurrentBranch(dir string) (string, error) {
	// e.g.: git rev-parse --abrev-ref HEAD
	return gitCommand(
		dir,
		"error trying to get current git branch name",
		"rev-parse", "--abbrev-ref", "HEAD",
	)
}

// VersionAtBranch returns the version file contents from the specific branch.
func VersionAtBranch(dir string, branchName string, versionFile string) (string, error) {
	// e.g.: git --no-pager show main:VERSION
	return gitCommand(
		dir,
		fmt.Sprintf("error trying to read %s from %s", versionFile, branchName),
		"--no-pager", "show", fmt.Sprintf("%s:%s", branchName, versionFile),
	)
}

func gitCommand(dir string, errMsg string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir

	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	if err := cmd.Run(); err != nil {
		return "", errors.Wrapf(err, "%s: %s", errMsg, stdErr.String())
	}

	return stdOut.String(), nil
}
