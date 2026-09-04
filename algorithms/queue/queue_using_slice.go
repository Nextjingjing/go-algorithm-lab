package queue

// QueueUsingSlice is a FIFO queue backed by a slice.
//
// Dequeue and Peek have a precondition that the queue is not empty.
type QueueUsingSlice struct {
	items []int
}

// NewQueueUsingSlice returns an empty QueueUsingSlice.
func NewQueueUsingSlice() QueueUsingSlice {
	return QueueUsingSlice{make([]int, 0)}
}

// Enqueue adds x to the back of the queue.
func (q *QueueUsingSlice) Enqueue(x int) {
	q.items = append(q.items, x)
}

// Dequeue removes and returns the value at the front of the queue.
func (q *QueueUsingSlice) Dequeue() int {
	if len(q.items) == 0 {
		panic("Queue is empty.")
	}
	first := q.items[0]
	q.items = q.items[1:len(q.items)]
	return first
}

// Peek returns the value at the front of the queue without removing it.
func (q *QueueUsingSlice) Peek() int {
	if len(q.items) == 0 {
		panic("Queue is empty.")
	}
	first := q.items[0]
	return first
}

// Empty reports whether the queue contains no values.
func (q *QueueUsingSlice) Empty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}
