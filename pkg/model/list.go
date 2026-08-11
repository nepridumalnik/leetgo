package model

import "slices"

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func NewList[T any](vals ...T) *ListNode[T] {
	var head *ListNode[T]

	for _, val := range slices.Backward(vals) {
		newNode := &ListNode[T]{Val: val}
		newNode.Next = head
		head = newNode
	}

	return head
}
