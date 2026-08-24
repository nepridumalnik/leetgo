package longestunequaladjacentgroupssubsequencei_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/longest_unequal_adjacent_groups_subsequence_i"

	"github.com/stretchr/testify/require"
)

type Input struct {
	Words  []string
	Groups []int
}

type testCase struct {
	Skip     bool
	Input    Input
	Expected []string
}

func Test_LongestUnequalAdjacentGroupsSubsequenceI(t *testing.T) {
	tests := []testCase{
		{
			Input: Input{
				Words:  []string{"e", "a", "b"},
				Groups: []int{0, 0, 1},
			},
			Expected: []string{"e", "b"},
		},
		{
			Input: Input{
				Words:  []string{"a", "b", "c", "d"},
				Groups: []int{1, 0, 1, 1},
			},
			Expected: []string{"a", "b", "c"},
		},
		{
			Input: Input{
				Words:  []string{"a", "b", "c", "d", "e"},
				Groups: []int{1, 0, 1, 0, 0},
			},
			Expected: []string{"a", "b", "c", "d"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := longestunequaladjacentgroupssubsequencei.GetLongestSubsequence(test.Input.Words, test.Input.Groups)
			require.Equal(t, test.Expected, result)
		})
	}
}
