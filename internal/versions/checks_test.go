package versions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaffenden/check-version/internal/semantic"
	"github.com/thaffenden/check-version/internal/versions"
)

func TestIsValidMajor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		was      semantic.Version
		now      semantic.Version
		expected bool
	}{
		"ReturnsTrueForValidMajor": {
			was:      semantic.Version{Major: 0, Minor: 0, Patch: 420},
			now:      semantic.Version{Major: 1, Minor: 0, Patch: 0},
			expected: true,
		},
		"ReturnsFalseWhenMajorTooHigh": {
			was:      semantic.Version{Major: 8, Minor: 3, Patch: 19},
			now:      semantic.Version{Major: 10, Minor: 0, Patch: 0},
			expected: false,
		},
		"ReturnsFalseWhenMinorIsNotReset": {
			was:      semantic.Version{Major: 2, Minor: 0, Patch: 4},
			now:      semantic.Version{Major: 3, Minor: 1, Patch: 0},
			expected: false,
		},
		"ReturnsFalseWhenPatchIsNotReset": {
			was:      semantic.Version{Major: 30, Minor: 812, Patch: 1},
			now:      semantic.Version{Major: 31, Minor: 0, Patch: 1},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := versions.IsValidMajor(tc.was, tc.now)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsValidMinor(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		was      semantic.Version
		now      semantic.Version
		expected bool
	}{
		"ReturnsTrueForValidMinor": {
			was:      semantic.Version{Major: 19, Minor: 4, Patch: 23},
			now:      semantic.Version{Major: 19, Minor: 5, Patch: 0},
			expected: true,
		},
		"ReturnsFalseWhenMinorTooHigh": {
			was:      semantic.Version{Major: 1, Minor: 4, Patch: 8},
			now:      semantic.Version{Major: 1, Minor: 6, Patch: 0},
			expected: false,
		},
		"ReturnsFalseWhenMajorIsIncreased": {
			was:      semantic.Version{Major: 7, Minor: 1, Patch: 9573},
			now:      semantic.Version{Major: 8, Minor: 2, Patch: 0},
			expected: false,
		},
		"ReturnsFalseWhenPatchIsNotReset": {
			was:      semantic.Version{Major: 365, Minor: 19, Patch: 4},
			now:      semantic.Version{Major: 365, Minor: 20, Patch: 4},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := versions.IsValidMinor(tc.was, tc.now)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsValidPatch(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		was      semantic.Version
		now      semantic.Version
		expected bool
	}{
		"ReturnsTrueForValidPatch": {
			was:      semantic.Version{Major: 1, Minor: 0, Patch: 4},
			now:      semantic.Version{Major: 1, Minor: 0, Patch: 5},
			expected: true,
		},
		"ReturnsFalseWhenPatchTooHigh": {
			was:      semantic.Version{Major: 0, Minor: 1, Patch: 4},
			now:      semantic.Version{Major: 0, Minor: 1, Patch: 6},
			expected: false,
		},
		"ReturnsFalseWhenMajorIsIncreased": {
			was:      semantic.Version{Major: 0, Minor: 1, Patch: 4},
			now:      semantic.Version{Major: 1, Minor: 1, Patch: 5},
			expected: false,
		},
		"ReturnsFalseWhenMinorIsIncreased": {
			was:      semantic.Version{Major: 0, Minor: 1, Patch: 4},
			now:      semantic.Version{Major: 0, Minor: 2, Patch: 5},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := versions.IsValidPatch(tc.was, tc.now)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
