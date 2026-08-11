package model_test

import (
	"fmt"
	"testing"

	"github.com/nepridumalnik/leetgo/pkg/model"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	Root     *model.TreeNode[int]
	Expected *model.TreeNode[int]
	Fail     bool
}

type testCases = []testCase

func Test_TreeNode_Creation(t *testing.T) {
	tests := testCases{
		// Simple tests
		{
			Root:     nil,
			Expected: nil,
		},
		{
			Root:     &model.TreeNode[int]{Val: 5},
			Expected: &model.TreeNode[int]{Val: 5},
		},
		{
			Root:     &model.TreeNode[int]{Val: 5},
			Expected: &model.TreeNode[int]{Val: 0},
			Fail:     true,
		},

		// Complicated tests
		{
			Root:     model.NewIntTreeNode(5),
			Expected: &model.TreeNode[int]{Val: 5},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("example_%d", i), func(t *testing.T) {
			t.Logf("Got %v. expected %v", test.Root, test.Expected)

			if !test.Fail {
				require.Equal(t, test.Root, test.Expected)
			} else {
				require.NotEqual(t, test.Root, test.Expected)
			}
		})
	}
}
