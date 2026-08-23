package numberof1bits_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/number_of_1_bits"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected int
}

func Test_NumberOf1Bits(t *testing.T) {
	tests := []testCase{
		{
			Input:    11,
			Expected: 3,
		},
		{
			Input:    128,
			Expected: 1,
		},
		{
			Input:    2147483645,
			Expected: 30,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}
		})

		result := numberof1bits.HammingWeight(test.Input)
		require.Equal(t, test.Expected, result)
	}
}
