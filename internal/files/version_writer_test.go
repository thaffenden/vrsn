package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/vrsn/internal/files"
	"github.com/thaffenden/vrsn/internal/test"
)

func TestWriteVersionToFile(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		parentDir   string
		inputFile   string
		newVersion  string
		assertError require.ErrorAssertionFunc
	}{
		"ReturnsErrorForUnsupportedVersionFile": {
			parentDir:   "bump",
			inputFile:   "foo.txt",
			newVersion:  "",
			assertError: require.Error,
		},
		"WritesVersionToCargoTOML": {
			parentDir:   "bump",
			inputFile:   "Cargo.toml",
			newVersion:  "2.14.741",
			assertError: require.NoError,
		},
		"ReturnsErrorForInvalidCargoTOML": {
			parentDir:   "no-version",
			inputFile:   "Cargo.toml",
			newVersion:  "",
			assertError: test.IsSentinelError(files.ErrGettingVersionFromTOML),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dir := filepath.Join("testdata", tc.parentDir)
			err := files.WriteVersionToFile(dir, tc.inputFile, tc.newVersion)
			tc.assertError(t, err)

			// Only assert the written contents if the writer func does not error.
			if err != nil {
				return
			}

			expected, err := os.ReadFile(filepath.Join("testdata", "all", tc.inputFile))
			require.NoError(t, err)

			actual, err := os.ReadFile(filepath.Clean(filepath.Join(dir, tc.inputFile)))
			require.NoError(t, err)

			assert.Equal(t, string(expected), string(actual))
		})
	}
}
