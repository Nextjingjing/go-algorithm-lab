package queue_test

import (
	"testing"

	queue "github.com/Nextjingjing/go-algorithm-lab/algorithms/queue"
)

func TestQueueUsingSlice(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		wantFront  int
		wantRemove int
	}{
		{
			name:       "removes the earliest value",
			values:     []int{10, 20, 30},
			wantFront:  10,
			wantRemove: 10,
		},
		{
			name:       "keeps duplicate values in order",
			values:     []int{7, 7},
			wantFront:  7,
			wantRemove: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := queue.NewQueueUsingSlice()
			if !q.Empty() {
				t.Fatal("new queue should be empty")
			}

			for _, value := range tt.values {
				q.Enqueue(value)
			}

			if got := q.Peek(); got != tt.wantFront {
				t.Fatalf("Peek() = %d, want %d", got, tt.wantFront)
			}
			if got := q.Dequeue(); got != tt.wantRemove {
				t.Fatalf("Dequeue() = %d, want %d", got, tt.wantRemove)
			}
			if q.Empty() {
				t.Fatal("queue should contain remaining values")
			}
		})
	}
}
