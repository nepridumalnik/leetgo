// https://leetcode.com/problems/reverse-linked-list
package reverselinkedlist

import "github.com/nepridumalnik/leetgo/pkg/model"

type ListNode = model.ListNode[int]

func ReverseList(head *ListNode) *ListNode {
	var (
		previous *ListNode
		current  *ListNode
	)

	for head != nil {
		current = head
		head = head.Next
		current.Next = previous
		previous = current
	}

	return current
}
