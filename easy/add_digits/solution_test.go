package adddigits_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/add_digits"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected int
}

func Test_AddDigits(t *testing.T) {
	tests := []testCase{
		{
			Input:    38, // 38 -> 3 + 8 -> 11, 11 = 1 + 1 -> 2
			Expected: 2,
		},
		{
			Input:    0,
			Expected: 0,
		},
		{
			Input:    1111,
			Expected: 4,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := adddigits.AddDigits(test.Input)
			require.Equal(t, result, test.Expected)
		})
	}
}
