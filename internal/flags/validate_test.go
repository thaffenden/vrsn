package flags_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thaffenden/check-version/internal/flags"
	"github.com/thaffenden/check-version/internal/test"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		was         string
		now         string
		assertError require.ErrorAssertionFunc
	}{
		"ReturnsErrorWhenNoWasOrNowSupplied": {
			was:         "",
			now:         "",
			assertError: test.IsSentinelError(flags.ErrNoValues),
		},
		"ReturnsErrorWhenNoWasSupplied": {
			was:         "",
			now:         "1.0.0",
			assertError: test.IsSentinelError(flags.ErrNoWasValue),
		},
		"ReturnsErrorWhenNoNowSupplied": {
			was:         "2.4.6",
			now:         "",
			assertError: test.IsSentinelError(flags.ErrNoNowValue),
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := flags.Validate(tc.was, tc.now)
			tc.assertError(t, err)
		})
	}
}
