package longestincreasingsubsequence_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/medium/longest_increasing_subsequence"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    []int
	Expected int
}

func Test_LongestIncreasingSubsequence(t *testing.T) {
	tests := []testCase{
		{
			Input:    []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			Expected: 6,
		},
		{
			Input:    []int{7, 7, 7, 7, 7, 7},
			Expected: 1,
		},
		{
			Input:    []int{4, 10, 4, 3, 8, 9},
			Expected: 3,
		},
		{
			Input:    []int{10, 9, 2, 5, 3, 7, 101, 18},
			Expected: 4,
		},
		{
			Input:    []int{0, 1, 0, 3, 2, 3},
			Expected: 4,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := longestincreasingsubsequence.LengthOfLIS(test.Input)
			require.Equal(t, test.Expected, result)
		})
	}
}
