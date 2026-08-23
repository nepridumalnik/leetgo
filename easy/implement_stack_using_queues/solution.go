// https://leetcode.com/problems/implement-stack-using-queues
package implementstackusingqueues

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0, 10),
	}
}

func (s *MyStack) Push(x int) {
	s.queue = append(s.queue, x)
	for i := 0; i < len(s.queue)-1; i++ {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *MyStack) Pop() int {
	x := s.queue[0]
	s.queue = s.queue[1:]
	return x
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}
