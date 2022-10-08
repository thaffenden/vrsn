// Package semantic holds logic for interacting with and validating a semantic
// version.
package semantic

import "fmt"

// Version holds the details of the semantic version parts.
type Version struct {
	Major int
	Minor int
	Patch int
}

// AsString formats the Version as string for ease of user output.
func (v Version) AsString() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
