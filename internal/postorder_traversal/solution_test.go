package postordertraversal_test

import (
	"fmt"
	"testing"

	postordertraversal "github.com/nepridumalnik/leetgo/internal/postorder_traversal"
	"github.com/nepridumalnik/leetgo/pkg/model"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	Skip     bool
	Root     *model.TreeNode[int]
	Expected []int
}

type testCases = []testCase

func TestSuite(t *testing.T) {
	tests := testCases{}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("testcase marked as skipped")
			}

			result := postordertraversal.PostorderTraversal(test.Root)
			require.Equal(t, result, test.Expected, "got %v, expected %v", result, test.Expected)
		})
	}
}
