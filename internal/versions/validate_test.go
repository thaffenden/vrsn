package versions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/check-version/internal/test"
	"github.com/thaffenden/check-version/internal/versions"
)

func TestValidateVersion(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         string
		errorExpected require.ErrorAssertionFunc
		expected      versions.SemVer
	}{
		"ReturnsVersionStructForValidInput": {
			input:         "34.9.154",
			errorExpected: require.NoError,
			expected: versions.SemVer{
				Major: 34,
				Minor: 9,
				Patch: 154,
			},
		},
		"ReturnsErrorIfVersionDoesNotContainSeparator": {
			input:         "100",
			errorExpected: test.IsSentinelError(versions.ErrNoVersionParts),
			expected:      versions.SemVer{},
		},
		"ReturnsErrorIfInputDoesNotHaveThreeParts": {
			input:         "2.2",
			errorExpected: test.IsSentinelError(versions.ErrNumVersionParts),
			expected:      versions.SemVer{},
		},
		"ReturnsErrorIfMajorVersionCannotBeConvertedToInt": {
			input:         "x.1.1",
			errorExpected: test.IsSentinelError(versions.ErrConvertingToInt),
			expected:      versions.SemVer{},
		},
		"ReturnsErrorIfMinorVersionCannotBeConvertedToInt": {
			input:         "1.x.1",
			errorExpected: test.IsSentinelError(versions.ErrConvertingToInt),
			expected:      versions.SemVer{},
		},
		"ReturnsErrorIfPatchVersionCannotBeConvertedToInt": {
			input:         "1.5.x",
			errorExpected: test.IsSentinelError(versions.ErrConvertingToInt),
			expected:      versions.SemVer{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := versions.Validate(tc.input)
			tc.errorExpected(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
