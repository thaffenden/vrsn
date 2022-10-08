package compare_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/check-version/internal/compare"
	"github.com/thaffenden/check-version/internal/test"
)

func TestVersions(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		was         string
		now         string
		assertError require.ErrorAssertionFunc
		expected    compare.ChangeType
	}{
		"ReturnsNoIncrementWhenVersionsAreTheSame": {
			was:         "1.0.0",
			now:         "1.0.0",
			assertError: require.NoError,
			expected:    compare.NoIncrement,
		},
		"ReturnsErrorForUnhandledScenario": {
			was:         "y.e.h",
			now:         "w.o.w",
			assertError: test.IsSentinelError(compare.ErrComparingVersions),
			expected:    0,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := compare.Versions(tc.was, tc.now)
			tc.assertError(t, err)

			assert.Equal(t, tc.expected.Message(), actual.Message())
		})
	}
}
