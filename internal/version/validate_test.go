package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/check-version/internal/semantic"
	"github.com/thaffenden/check-version/internal/test"
	"github.com/thaffenden/check-version/internal/version"
)

func TestValidateVersion(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         string
		errorExpected require.ErrorAssertionFunc
		expected      semantic.Version
	}{
		"ReturnsVersionStructForValidInput": {
			input:         "34.9.154",
			errorExpected: require.NoError,
			expected: semantic.Version{
				Major: 34,
				Minor: 9,
				Patch: 154,
			},
		},
		"ReturnsErrorIfVersionDoesNotContainSeparator": {
			input:         "100",
			errorExpected: test.IsSentinelError(version.ErrNoVersionParts),
			expected:      semantic.Version{},
		},
		"ReturnsErrorIfInputDoesNotHaveThreeParts": {
			input:         "2.2",
			errorExpected: test.IsSentinelError(version.ErrNumVersionParts),
			expected:      semantic.Version{},
		},
		"ReturnsErrorIfMajorVersionCannotBeConvertedToInt": {
			input:         "x.1.1",
			errorExpected: test.IsSentinelError(version.ErrConvertingToInt),
			expected:      semantic.Version{},
		},
		"ReturnsErrorIfMinorVersionCannotBeConvertedToInt": {
			input:         "1.x.1",
			errorExpected: test.IsSentinelError(version.ErrConvertingToInt),
			expected:      semantic.Version{},
		},
		"ReturnsErrorIfPatchVersionCannotBeConvertedToInt": {
			input:         "1.5.x",
			errorExpected: test.IsSentinelError(version.ErrConvertingToInt),
			expected:      semantic.Version{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := version.Validate(tc.input)
			tc.errorExpected(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}