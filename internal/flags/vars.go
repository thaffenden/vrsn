// Package flags holds logics for use of CLI flags.
package flags

var (
	// BaseBranch is the name of the base branch used when auto detecting version
	// file changes.
	BaseBranch string
	// Now is the variable for the CLI flag --now.
	Now string
	// Was is the variable for the CLI flag --was.
	Was string
)
