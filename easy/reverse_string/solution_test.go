package reversestring_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/reverse_string"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Input    []byte
	Expected []byte
}

func Test_ReverseString(t *testing.T) {
	tests := []testCase{
		{
			Input:    []byte{'h', 'e', 'l', 'l', 'o'},
			Expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			reversestring.ReverseString(test.Input)
			require.Equal(t, test.Expected, test.Input)
		})
	}
}
