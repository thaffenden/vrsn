package files_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/vrsn/internal/files"
)

func TestGetVersionFromFile(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputFile   string
		assertError require.ErrorAssertionFunc
		expected    string
	}{
		"ReturnsErrorForUnsupportedVersionFile": {
			inputFile:   "foo.txt",
			assertError: require.Error,
			expected:    "",
		},
		"ReturnsVersionFromVERSIONFile": {
			inputFile:   "VERSION",
			assertError: require.NoError,
			expected:    "6.6.6",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dir := filepath.Join("testdata", "all")
			actual, err := files.GetVersionFromFile(dir, tc.inputFile)
			tc.assertError(t, err)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
