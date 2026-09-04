package stack

// StackUsingQueues is a LIFO stack implemented using queue operations.
//
// Pop and Top have a precondition that the stack is not empty.
type StackUsingQueues struct {
	queue []int
}

// NewStackUsingQueues returns an empty StackUsingQueues.
func NewStackUsingQueues() StackUsingQueues {
	return StackUsingQueues{make([]int, 0)}
}

// Push adds x to the top of the stack.
func (s *StackUsingQueues) Push(x int) {
	s.queue = append(s.queue, x)
}

// Pop removes and returns the value at the top of the stack.
func (s *StackUsingQueues) Pop() int {
	if s.Empty() {
		panic("Empty")
	}

	for i := 0; i < len(s.queue)-1; i++ {
		first := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, first)
	}
	first := s.queue[0]
	s.queue = s.queue[1:]
	return first
}

// Top returns the value at the top of the stack without removing it.
func (s *StackUsingQueues) Top() int {
	if s.Empty() {
		panic("Empty")
	}
	first := 0
	for i := 0; i < len(s.queue); i++ {
		first = s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, first)
	}
	return first
}

// Empty reports whether the stack contains no values.
func (s *StackUsingQueues) Empty() bool {
	if len(s.queue) == 0 {
		return true
	}
	return false
}
