package stack_test

import (
	"testing"

	stack "github.com/Nextjingjing/go-algorithm-lab/algorithms/stack"
)

func TestStackUsingQueues(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		wantTop    int
		wantRemove int
	}{
		{
			name:       "removes the most recently added value",
			values:     []int{10, 20, 30},
			wantTop:    30,
			wantRemove: 30,
		},
		{
			name:       "keeps duplicate values",
			values:     []int{7, 7},
			wantTop:    7,
			wantRemove: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack.NewStackUsingQueues()
			if !s.Empty() {
				t.Fatal("new stack should be empty")
			}

			for _, value := range tt.values {
				s.Push(value)
			}

			if got := s.Top(); got != tt.wantTop {
				t.Fatalf("Top() = %d, want %d", got, tt.wantTop)
			}
			if got := s.Pop(); got != tt.wantRemove {
				t.Fatalf("Pop() = %d, want %d", got, tt.wantRemove)
			}
			if s.Empty() {
				t.Fatal("stack should contain remaining values")
			}
		})
	}
}
