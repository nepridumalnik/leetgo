package greatestcommondivisorofstrings_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/greatest_common_divisor_of_strings"

	"github.com/stretchr/testify/require"
)

type testSuiteGreatestCommonDivisorOfString struct {
	InputA   string
	InputB   string
	Expected string
	Skip     bool
}

func Test_GreatestCommonDivisorOfString(t *testing.T) {
	tests := []testSuiteGreatestCommonDivisorOfString{
		{
			InputA:   "ABABABAB",
			InputB:   "ABAB",
			Expected: "ABAB",
		},
		{
			InputA:   "TAUXXTAUXXTAUXXTAUXXTAUXX",
			InputB:   "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			Expected: "TAUXX",
		},
		{
			InputA:   "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX",
			InputB:   "TAUXXTAUXXTAUXXTAUXXTAUXX",
			Expected: "TAUXX",
		},
		{
			InputA:   "EKEKE",
			InputB:   "KEK",
			Expected: "",
		},
		{
			InputA:   "ARARARAR",
			InputB:   "ARARAR",
			Expected: "AR",
		},
		{
			InputA:   "WOWWOWWOWWOW",
			InputB:   "WOW",
			Expected: "WOW",
		},
		{
			InputA:   "ABCABC",
			InputB:   "ABC",
			Expected: "ABC",
		},
		{
			InputA:   "ABABAB",
			InputB:   "ABAB",
			Expected: "AB",
		},
		{
			InputA:   "LEET",
			InputB:   "CODE",
			Expected: "",
		},
		{
			InputA:   "AAAAAB",
			InputB:   "AAA",
			Expected: "",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip()
			}

			result := greatestcommondivisorofstrings.GcdOfStrings(test.InputA, test.InputB)
			require.Equal(t, result, test.Expected, "got %s, expected %s", result, test.Expected)
		})
	}
}
