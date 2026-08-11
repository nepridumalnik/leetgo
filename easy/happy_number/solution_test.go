package happynumber_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/happy_number"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected bool
}

func Test_HappyNumber(t *testing.T) {
	tests := []testCase{
		{
			Input:    19,
			Expected: true,
		},
		{
			Input:    2,
			Expected: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := happynumber.IsHappy(test.Input)
			require.Equal(t, result, test.Expected)
		})
	}
}
