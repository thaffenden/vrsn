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
			was:      semantic.Version{Major: 0},
			now:      semantic.Version{Major: 1},
			expected: true,
		},
		"ReturnsFalseForInValidMajor": {
			was:      semantic.Version{Major: 9},
			now:      semantic.Version{Major: 8},
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
			was:      semantic.Version{Minor: 4},
			now:      semantic.Version{Minor: 5},
			expected: true,
		},
		"ReturnsFalseForInValidMinor": {
			was:      semantic.Version{Minor: 4},
			now:      semantic.Version{Minor: 6},
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
			was:      semantic.Version{Patch: 4},
			now:      semantic.Version{Patch: 5},
			expected: true,
		},
		"ReturnsFalseForInValidPatch": {
			was:      semantic.Version{Patch: 4},
			now:      semantic.Version{Patch: 6},
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
