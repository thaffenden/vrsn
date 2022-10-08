package versions_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thaffenden/check-version/internal/test"
	"github.com/thaffenden/check-version/internal/versions"
)

func TestCompare(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		was         string
		now         string
		assertError require.ErrorAssertionFunc
	}{
		"ReturnsVersionNotBumpedErrorWhenVersionsAreTheSame": {
			was:         "1.0.0",
			now:         "1.0.0",
			assertError: test.IsSentinelError(versions.ErrVersionNotBumped),
		},
		"ReturnsErrorWhenWasFailsValidation": {
			was:         "",
			now:         "1.1.1",
			assertError: require.Error,
		},
		"ReturnsErrorWhenNowFailsValidation": {
			was:         "1.1.1",
			now:         "",
			assertError: require.Error,
		},
		"ReturnsInvalidBumpErrorWhenNotValidSemVer": {
			was:         "1.0.0",
			now:         "1.0.3",
			assertError: test.IsSentinelError(versions.ErrInvalidBump),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := versions.Compare(tc.was, tc.now)
			tc.assertError(t, err)
		})
	}
}