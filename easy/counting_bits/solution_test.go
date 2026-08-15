package countingbits_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/counting_bits"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected []int
}

func Test_CountingBits(t *testing.T) {
	tests := []testCase{
		{
			Input:    16,
			Expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1},
		},
		{
			Input:    8,
			Expected: []int{0, 1, 1, 2, 1, 2, 2, 3, 1},
		},
		{
			Input:    2,
			Expected: []int{0, 1, 1},
		},
		{
			Input:    5,
			Expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := countingbits.CountBits(test.Input)
			require.Equal(t, test.Expected, result)
		})
	}
}
