package hashmap_test

import (
	"slices"
	"testing"

	hashmap "github.com/Nextjingjing/go-algorithm-lab/algorithms/hash-map"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "found pair",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "pair uses duplicate values at different indexes",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "no pair",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashmap.TwoSum(tt.nums, tt.target)

			if !slices.Equal(got, tt.want) {
				t.Fatalf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
