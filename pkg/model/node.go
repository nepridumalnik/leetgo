package model

import "slices"

type TreeNode[T any] struct {
	Val   int
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewIntTreeNode(x ...int) *TreeNode[int] {
	slices.Sort(x)
	var root *TreeNode[int]

	return root
}
