package files_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/vrsn/internal/files"
)

func TestWriteVersionToFile(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		parentDir   string
		inputFile   string
		newVersion  string
		assertError require.ErrorAssertionFunc
	}{
		"": {},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			dir := filepath.Join("testdata", "bump")
			err := files.WriteVersionToFile(dir, tc.inputFile, tc.newVersion)
			tc.assertError(t, err)

			expected, err := ioutil.ReadFile(filepath.Join("testdata", "all", tc.inputFile))
			tc.assertError(t, err)

			actual, err := ioutil.ReadFile(filepath.Join("testdata", "bump", tc.inputFile))
			tc.assertError(t, err)

			assert.Equal(t, expected, actual)
		})
	}
}
