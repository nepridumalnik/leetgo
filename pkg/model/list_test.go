package model_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/pkg/model"
	"github.com/stretchr/testify/require"
)

type testCaseList struct {
	Input    []int
	Expected *model.ListNode[int]
}

func Test_ListCreation(t *testing.T) {
	tests := []testCaseList{
		{
			Input: []int{1, 2, 3, 4, 5},
			Expected: &model.ListNode[int]{
				Val: 1, Next: &model.ListNode[int]{
					Val: 2, Next: &model.ListNode[int]{
						Val: 3, Next: &model.ListNode[int]{
							Val: 4, Next: &model.ListNode[int]{
								Val: 5}}}}},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			result := model.NewList(test.Input...)
			require.Equal(t, result, test.Expected)
		})
	}

}
