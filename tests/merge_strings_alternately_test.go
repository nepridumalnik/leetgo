package tests

import (
	"fmt"
	"testing"

	mergestringsalternately "github.com/nepridumalnik/leetgo/internal/merge_strings_alternately"

	"github.com/stretchr/testify/require"
)

type testSuiteMergeStringsAlternately struct {
	InputA   string
	InputB   string
	Expected string
	Skip     bool
}

func Test_MergeStringsAlternately(t *testing.T) {
	tests := []testSuiteMergeStringsAlternately{
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

			result := mergestringsalternately.MergeAlternately(test.InputA, test.InputB)
			require.Equal(t, result, test.Expected, "got %s, expected %s", result, test.Expected)
		})
	}
}
