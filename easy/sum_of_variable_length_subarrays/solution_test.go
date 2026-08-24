package sumofvariablelengthsubarrays_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/sum_of_variable_length_subarrays"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    []int
	Expected int
}

func Test_SumOfVariableLengthSubarrays(t *testing.T) {
	tests := []testCase{
		{
			Input:    []int{2, 3, 1},
			Expected: 11,
		},
		{
			Input:    []int{3, 1, 1, 2},
			Expected: 13,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := sumofvariablelengthsubarrays.SubarraySum(test.Input)
			require.Equal(t, test.Expected, result)
		})
	}
}
