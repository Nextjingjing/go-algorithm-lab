package queue

// MyQueue is a FIFO queue implemented using stacks.
//
// Pop and Peek have a precondition that the queue is not empty.
type MyQueue struct {
	stack []int
	queue []int
}

// Constructor returns an empty queue.
func Constructor() MyQueue {
	return MyQueue{
		stack: make([]int, 0),
		queue: make([]int, 0),
	}
}

// Push adds x to the back of the queue.
func (q *MyQueue) Push(x int) {
	q.stack = append(q.stack, x)
}

// Pop removes and returns the value at the front of the queue.
func (q *MyQueue) Pop() int {
	if len(q.queue) == 0 {
		for i := len(q.stack) - 1; i >= 0; i-- {
			q.queue = append(q.queue, q.stack[len(q.stack)-1])
			q.stack = q.stack[:len(q.stack)-1]
		}
	}

	top := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]
	return top
}

// Peek returns the value at the front of the queue.
func (q *MyQueue) Peek() int {
	if len(q.queue) > 0 {
		return q.queue[len(q.queue)-1]
	}
	for i := len(q.stack) - 1; i >= 0; i-- {
		q.queue = append(q.queue, q.stack[len(q.stack)-1])
		q.stack = q.stack[:len(q.stack)-1]
	}
	return q.queue[len(q.queue)-1]
}

// Empty reports whether the queue contains no values.
func (q *MyQueue) Empty() bool {
	if len(q.queue) == 0 && len(q.stack) == 0 {
		return true
	}
	return false
}
