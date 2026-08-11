package excelsheetcolumntitle_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/internal/missing_number"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	Input    []int
	Expected int
	Skip     bool
}

func Test_MissingNumber(t *testing.T) {
	tests := []testCase{
		{
			Input:    []int{3, 0, 1},
			Expected: 2,
		},
		{
			Input:    []int{0, 1},
			Expected: 2,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("testcase marked as skipped")
			}
			result := excelsheetcolumntitle.MissingNumber(test.Input)
			require.Equal(t, result, test.Expected)
		})
	}
}
