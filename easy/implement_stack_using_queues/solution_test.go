package implementstackusingqueues_test

import (
	"testing"

	"github.com/nepridumalnik/leetgo/easy/implement_stack_using_queues"

	"github.com/stretchr/testify/require"
)

func Test_ImplementStackUsingQueues(t *testing.T) {
	stack := implementstackusingqueues.Constructor()

	stack.Push(1)
	stack.Push(2)

	top := stack.Top()
	require.Equal(t, 2, top)

	pop := stack.Pop()
	require.Equal(t, 2, pop)

	empty := stack.Empty()
	require.Equal(t, false, empty)
}
