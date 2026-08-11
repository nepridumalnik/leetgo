package reverselinkedlist_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/easy/reverse_linked_list"
	"github.com/nepridumalnik/leetgo/pkg/model"

	"github.com/stretchr/testify/require"
)

type ListNode = model.ListNode[int]

type testCase struct {
	Skip     bool
	Input    *ListNode
	Expected *ListNode
}

func Test_ReverseLinkedList(t *testing.T) {
	tests := []testCase{
		{
			Input:    model.NewList(1, 2, 1),
			Expected: model.NewList(1, 2, 1),
		},
		{
			Input:    model.NewList(1, 2, 3),
			Expected: model.NewList(3, 2, 1),
		},
		{
			Input:    model.NewList[int](),
			Expected: model.NewList[int](),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			if test.Skip {
				t.Skip("skipped")
			}

			result := reverselinkedlist.ReverseList(test.Input)
			require.Equal(t, result, test.Expected)
		})
	}
}
