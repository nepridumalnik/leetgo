package complementofbase10integer_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/complement_of_base_10_integer"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected int
}

func Test_ComplementOfBase10Integer(t *testing.T) {
	tests := []testCase{
		{
			Input:    5,
			Expected: 2,
		},
		{
			Input:    7,
			Expected: 0,
		},
		{
			Input:    10,
			Expected: 5,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := complementofbase10integer.BitwiseComplement(test.Input)
			require.Equal(t, test.Expected, result)
		})
	}
}
