package longestsubstringwithoutrepeatingcharacters_test

import (
	"fmt"
	"testing"

	longestsubstringwithoutrepeatingcharacters "github.com/nepridumalnik/leetgo/medium/longest_substring_without_repeating_characters"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    string
	Expected int
}

func Test_LongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	tests := []testCase{
		{
			Input:    "cdcda",
			Expected: 3,
		},
		{
			Input:    "1R1T7",
			Expected: 4,
		},
		{
			Input:    "abbac",
			Expected: 3,
		},
		{
			Input:    "abcabcbb",
			Expected: 3,
		},
		{
			Input:    "bbb",
			Expected: 1,
		},
		{
			Input:    "pwwkew",
			Expected: 3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := longestsubstringwithoutrepeatingcharacters.LengthOfLongestSubstring(test.Input)
			require.Equal(t, test.Expected, result)
		})
	}
}
