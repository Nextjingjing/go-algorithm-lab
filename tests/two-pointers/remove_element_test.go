package twopointers_test

import (
	"slices"
	"testing"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		val    int
		wantK  int
		prefix []int
	}{
		{
			name:   "removes values from both ends",
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:    2,
			wantK:  5,
			prefix: []int{0, 1, 3, 0, 4},
		},
		{
			name:   "removes values at both sides",
			nums:   []int{3, 2, 2, 3},
			val:    3,
			wantK:  2,
			prefix: []int{2, 2},
		},
		{
			name:   "removes every value",
			nums:   []int{1, 1, 1},
			val:    1,
			wantK:  0,
			prefix: []int{},
		},
		{
			name:   "removes nothing",
			nums:   []int{1, 2, 3},
			val:    4,
			wantK:  3,
			prefix: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotK := twopointers.RemoveElement(tt.nums, tt.val)

			if gotK != tt.wantK {
				t.Fatalf("RemoveElement(%v, %d) returned k=%d, want %d", tt.nums, tt.val, gotK, tt.wantK)
			}

			if !slices.Equal(tt.nums[:gotK], tt.prefix) {
				t.Fatalf("RemoveElement(%v, %d) prefix = %v, want %v", tt.nums, tt.val, tt.nums[:gotK], tt.prefix)
			}
		})
	}
}
