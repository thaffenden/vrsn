package files_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thaffenden/vrsn/internal/files"
	"github.com/thaffenden/vrsn/internal/test"
)

func TestGetVersionFilesInDirectory(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		directory     string
		assertError   require.ErrorAssertionFunc
		expectedFiles []string
	}{
		"ReturnsErrorWhenNoVersionFilesFound": {
			directory:     "testdata/empty",
			assertError:   test.IsSentinelError(files.ErrNoVersionFilesInDir),
			expectedFiles: []string{},
		},
		"ReturnsSupportVersionFilesWhenFound": {
			directory:   "testdata/all",
			assertError: require.NoError,
			expectedFiles: []string{
				"Cargo.toml",
				"package.json",
				"pyproject.toml",
				"VERSION",
			},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			path := filepath.FromSlash(tc.directory)
			actual, err := files.GetVersionFilesInDirectory(path)
			tc.assertError(t, err)
			reflect.DeepEqual(tc.expectedFiles, actual)
		})
	}
}
