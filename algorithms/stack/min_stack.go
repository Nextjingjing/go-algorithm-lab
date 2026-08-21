package stack

// MinStack is a stack that supports retrieving its minimum value.
//
// Pop, Top, and GetMin have a precondition that the stack is not empty.
type MinStack struct {
	stack []int
	mins  []int
}

// Constructor returns an empty MinStack.
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mins:  []int{},
	}
}

// Push adds val to the top of the stack.
func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)

	if len(m.mins) > 0 {
		m.mins = append(m.mins, min(m.mins[len(m.mins)-1], val))
	} else {
		m.mins = append(m.mins, val)
	}

}

// Pop removes the value at the top of the stack.
func (m *MinStack) Pop() {
	if len(m.stack) == 0 {
		return
	}
	m.stack = m.stack[:len(m.stack)-1]
	m.mins = m.mins[:len(m.mins)-1]
}

// Top returns the value at the top of the stack.
func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

// GetMin returns the smallest value currently in the stack.
func (m *MinStack) GetMin() int {
	return m.mins[len(m.mins)-1]
}
