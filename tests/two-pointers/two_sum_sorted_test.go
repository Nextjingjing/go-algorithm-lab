package twopointers_test

import (
	"slices"
	"testing"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "pair at both ends",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "pair in the middle",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "negative and zero values",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "duplicate values form the pair",
			numbers: []int{1, 1, 3, 4},
			target:  2,
			want:    []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twopointers.TwoSumSorted(tt.numbers, tt.target)

			if !slices.Equal(got, tt.want) {
				t.Fatalf("TwoSumSorted(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
