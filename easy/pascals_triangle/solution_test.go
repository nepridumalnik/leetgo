package pascalstriangle_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/pascals_triangle"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    int
	Expected [][]int
}

func Test_PascalsTriangle(t *testing.T) {
	tests := []testCase{
		{
			Input:    1,
			Expected: [][]int{{1}},
			Skip:     true,
		},
		{
			Input:    2,
			Expected: [][]int{{1}, {1, 1}},
			Skip:     true,
		},
		{
			Input:    3,
			Expected: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			Input:    4,
			Expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			Input:    5,
			Expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := pascalstriangle.Generate(test.Input)
			require.Equal(t, result, test.Expected)
		})
	}
}
