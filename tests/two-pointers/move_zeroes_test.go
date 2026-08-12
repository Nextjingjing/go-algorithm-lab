package twopointers_test

import (
	"slices"
	"testing"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "moves zeros to the end", nums: []int{0, 1, 0, 3, 12}, want: []int{1, 3, 12, 0, 0}},
		{name: "keeps multiple leading zeros behind non-zero values", nums: []int{0, 0, 1}, want: []int{1, 0, 0}},
		{name: "leaves slice without zeros unchanged", nums: []int{1, 2, 3}, want: []int{1, 2, 3}},
		{name: "leaves empty slice unchanged", nums: []int{}, want: []int{}},
		{name: "preserves negative non-zero values", nums: []int{-1, 0, -2, 0, 3}, want: []int{-1, -2, 3, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			twopointers.MoveZeroes(tt.nums)

			if !slices.Equal(tt.nums, tt.want) {
				t.Fatalf("MoveZeroes() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
