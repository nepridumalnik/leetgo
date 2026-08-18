package maximumerasurevalue_test

import (
	"fmt"
	"testing"

	maximumerasurevalue "github.com/nepridumalnik/leetgo/medium/maximum_erasure_value"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    []int
	Expected int
}

func Test_MaximumErasureValue(t *testing.T) {
	tests := []testCase{
		{
			Input:    []int{10000, 1, 10000, 1, 1},
			Expected: 10001,
		},
		{
			Input:    []int{4, 2, 4, 5, 6},
			Expected: 17,
		},
		{
			Input:    []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			Expected: 8,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}
		})

		result := maximumerasurevalue.MaximumUniqueSubarray(test.Input)
		require.Equal(t, test.Expected, result)
	}
}
