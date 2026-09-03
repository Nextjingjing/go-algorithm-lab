package queue_test

import (
	"testing"

	queue "github.com/Nextjingjing/go-algorithm-lab/algorithms/queue"
)

func TestMyQueue(t *testing.T) {
	q := queue.Constructor()

	if !q.Empty() {
		t.Fatal("new queue should be empty")
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if got := q.Peek(); got != 1 {
		t.Fatalf("Peek() = %d, want 1", got)
	}
	if got := q.Pop(); got != 1 {
		t.Fatalf("Pop() = %d, want 1", got)
	}
	if got := q.Pop(); got != 2 {
		t.Fatalf("Pop() = %d, want 2", got)
	}
	if q.Empty() {
		t.Fatal("queue should not be empty after two of three values are removed")
	}
	if got := q.Peek(); got != 3 {
		t.Fatalf("Peek() = %d, want 3", got)
	}

	q.Pop()
	if !q.Empty() {
		t.Fatal("queue should be empty after removing all values")
	}
}
