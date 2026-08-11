package convertsortedarraytobinarysearchtree_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/convert_sorted_array_to_binary_search_tree"
	"github.com/nepridumalnik/leetgo/pkg/model"

	"github.com/stretchr/testify/require"
)

type TreeNode = model.TreeNode[int]

type testCase struct {
	Skip     bool
	Input    []int
	Expected *TreeNode
}

func Test_ConvertSortedArrayToBinarySearchTree(t *testing.T) {
	tests := []testCase{
		{
			Input:    []int{1, 2, 3},
			Expected: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		},
		{
			Input:    []int{-10, -3, 0, 5, 9},
			Expected: &TreeNode{Val: 0, Left: &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}},
		},
		{
			Input:    []int{1, 3},
			Expected: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := convertsortedarraytobinarysearchtree.SortedArrayToBST(test.Input)
			require.Equal(t, result, test.Expected)
		})
	}
}
