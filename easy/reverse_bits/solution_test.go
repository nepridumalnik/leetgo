package reversebits_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/reverse_bits"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected int
}

func Test_ReverseBits(t *testing.T) {
	tests := []testCase{
		{
			Input:    43261596,
			Expected: 964176192,
		},
		{
			Input:    2147483644,
			Expected: 1073741822,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := reversebits.ReverseBits(test.Input)
			require.Equal(t, test.Expected, result)
		})
	}
}
