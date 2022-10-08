package version_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/check-version/internal/compare"
	"github.com/thaffenden/check-version/internal/test"
)

func TestValidateVersion(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         string
		errorExpected require.ErrorAssertionFunc
		expected      compare.Version
	}{
		"ReturnsVersionStructForValidInput": {
			input:         "34.9.154",
			errorExpected: require.NoError,
			expected: compare.Version{
				Major: 34,
				Minor: 9,
				Patch: 154,
			},
		},
		"ReturnsErrorIfVersionDoesNotContainSeparator": {
			input:         "100",
			errorExpected: test.IsSentinelError(compare.ErrNoVersionParts),
			expected:      compare.Version{},
		},
		"ReturnsErrorIfInputDoesNotHaveThreeParts": {
			input:         "2.2",
			errorExpected: test.IsSentinelError(compare.ErrNumVersionParts),
			expected:      compare.Version{},
		},
		"ReturnsErrorIfMajorVersionCannotBeConvertedToInt": {
			input:         "x.1.1",
			errorExpected: test.IsSentinelError(compare.ErrConvertingToInt),
			expected:      compare.Version{},
		},
		"ReturnsErrorIfMinorVersionCannotBeConvertedToInt": {
			input:         "1.x.1",
			errorExpected: test.IsSentinelError(compare.ErrConvertingToInt),
			expected:      compare.Version{},
		},
		"ReturnsErrorIfPatchVersionCannotBeConvertedToInt": {
			input:         "1.5.x",
			errorExpected: test.IsSentinelError(compare.ErrConvertingToInt),
			expected:      compare.Version{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := compare.ValidateVersion(tc.input)
			tc.errorExpected(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
