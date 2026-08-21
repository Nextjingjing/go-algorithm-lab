package stack_test

import (
	"testing"

	stack "github.com/Nextjingjing/go-algorithm-lab/algorithms/stack"
)

func TestMinStack(t *testing.T) {
	tests := []struct {
		name       string
		pushValues []int
		popCount   int
		wantTop    int
		wantMin    int
	}{
		{
			name:       "minimum changes after pop",
			pushValues: []int{-2, 0, -3},
			popCount:   1,
			wantTop:    0,
			wantMin:    -2,
		},
		{
			name:       "minimum with duplicate values",
			pushValues: []int{2, 2, 1, 1},
			popCount:   1,
			wantTop:    1,
			wantMin:    1,
		},
		{
			name:       "minimum at bottom",
			pushValues: []int{-5, 4, 7},
			popCount:   0,
			wantTop:    7,
			wantMin:    -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			minStack := stack.Constructor()

			for _, value := range tt.pushValues {
				minStack.Push(value)
			}
			for i := 0; i < tt.popCount; i++ {
				minStack.Pop()
			}

			if got := minStack.Top(); got != tt.wantTop {
				t.Fatalf("Top() = %d, want %d", got, tt.wantTop)
			}
			if got := minStack.GetMin(); got != tt.wantMin {
				t.Fatalf("GetMin() = %d, want %d", got, tt.wantMin)
			}
		})
	}
}
