// Package semantic holds logic for interacting with and validating a semantic
// version.
package semantic

// Version holds the details of the semantic version parts.
type Version struct {
	Major int
	Minor int
	Patch int
}
