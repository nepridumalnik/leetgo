package issubsequence_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/is_subsequence"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	InputS   string
	InputT   string
	Expected bool
}

func Test_IsSubsequence(t *testing.T) {
	tests := []testCase{
		{
			InputS:   "abc",
			InputT:   "ahbgdc",
			Expected: true,
		},
		{
			InputS:   "axc",
			InputT:   "ahbgdc",
			Expected: false,
		},
		{
			InputS:   "",
			InputT:   "abc",
			Expected: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := issubsequence.IsSubsequence(test.InputS, test.InputT)
			require.Equal(t, result, test.Expected)
		})
	}
}
