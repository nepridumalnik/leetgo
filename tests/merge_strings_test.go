package merge_strings_alternately_test

import (
	"fmt"
	"testing"

	merge_strings_alternately "github.com/nepridumalnik/leetgo/tests"
	"github.com/stretchr/testify/require"
)

type testSuite struct {
	InputA   string
	InputB   string
	Expected string
	Skip     bool
}

func TestTable(t *testing.T) {
	tests := []testSuite{
		{
			InputA:   "cdf",
			InputB:   "a",
			Expected: "cadf",
		},
		{
			InputA:   "abc",
			InputB:   "pqr",
			Expected: "apbqcr",
		},
		{
			InputA:   "ab",
			InputB:   "pqrs",
			Expected: "apbqrs",
		},
		{
			InputA:   "abcd",
			InputB:   "pq",
			Expected: "apbqcd",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip()
			}

			result := merge_strings_alternately.MergeAlternately(test.InputA, test.InputB)
			require.Equal(t, result, test.Expected, "got %s, expected %s", result, test.Expected)
		})
	}
}
