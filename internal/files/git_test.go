package files_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/vrsn/internal/files"
)

func TestIsGitDir(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputDir      string
		errorExpected require.ErrorAssertionFunc
		expected      bool
	}{
		"ReturnsTrueIfIsGitDir": {
			inputDir:      "testdata/all",
			errorExpected: require.NoError,
			expected:      true,
		},
		"ReturnsFalseIfNotGitDir": {
			inputDir:      "testdata/no-version",
			errorExpected: require.NoError,
			expected:      false,
		},
		"ReturnsErrorIfDirectoryDoesNotExist": {
			inputDir:      "testdata/foo",
			errorExpected: require.Error,
			expected:      false,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := files.IsGitDir(tc.inputDir)
			tc.errorExpected(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
