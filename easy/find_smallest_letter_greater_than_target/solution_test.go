package findsmallestlettergreaterthantarget_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/find_smallest_letter_greater_than_target"

	"github.com/stretchr/testify/require"
)

type Input struct {
	Letters []byte
	Target  byte
}

type testCase struct {
	Skip     bool
	Input    Input
	Expected byte
}

func Test_FindSmallestLetterGreaterThanTarget(t *testing.T) {
	tests := []testCase{
		{
			Input: Input{
				Letters: []byte{'x', 'x', 'y', 'y'},
				Target:  'z',
			},
			Expected: 'x',
		},
		{
			Input: Input{
				Letters: []byte{'c', 'f', 'j'},
				Target:  'a',
			},
			Expected: 'c',
		},
		{
			Input: Input{
				Letters: []byte{'c', 'f', 'j'},
				Target:  'c',
			},
			Expected: 'f',
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := findsmallestlettergreaterthantarget.NextGreatestLetter(test.Input.Letters, test.Input.Target)
			require.Equal(t, rune(test.Expected), rune(result))
		})
	}
}
