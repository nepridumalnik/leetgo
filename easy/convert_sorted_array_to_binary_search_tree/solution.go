// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree
package convertsortedarraytobinarysearchtree

import "github.com/nepridumalnik/leetgo/pkg/model"

type TreeNode = model.TreeNode[int]

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	middle := (0 + len(nums)) / 2
	node := &TreeNode{Val: nums[middle]}

	node.Left = SortedArrayToBST(nums[:middle])
	node.Right = SortedArrayToBST(nums[middle+1:])

	return node
}
