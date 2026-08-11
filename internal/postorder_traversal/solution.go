// https://leetcode.com/problems/binary-tree-postorder-traversal/
package postordertraversal

import (
	"github.com/nepridumalnik/leetgo/pkg/model"
)

type TreeNode = model.TreeNode[int]

func PostorderTraversal(root *TreeNode) []int {
	result := []int{}
	var appendValue func(n *TreeNode)
	appendValue = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Right != nil {
			appendValue(n.Right)
			result = append(result, n.Right.Val)
		}
		if n.Left != nil {
			appendValue(n.Left)
			result = append(result, n.Left.Val)
		}
		result = append(result, n.Val)
	}

	return result
}
