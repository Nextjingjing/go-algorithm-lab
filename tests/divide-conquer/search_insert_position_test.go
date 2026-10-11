package divideconquer_test

import (
	"testing"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "finds first value", nums: []int{1, 3, 5, 6}, target: 1, want: 0},
		{name: "finds last value", nums: []int{1, 3, 5, 6}, target: 6, want: 3},
		{name: "finds only value", nums: []int{4}, target: 4, want: 0},
		{name: "inserts among negative values", nums: []int{-8, -3, 2}, target: -5, want: 1},
		{
			name:   "finds existing value",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "inserts between values",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "inserts after all values",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "inserts before all values",
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
		{
			name:   "inserts before one value",
			nums:   []int{4},
			target: -2,
			want:   0,
		},
		{
			name:   "inserts after one value",
			nums:   []int{4},
			target: 9,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divideconquer.SearchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Fatalf("SearchInsert(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
