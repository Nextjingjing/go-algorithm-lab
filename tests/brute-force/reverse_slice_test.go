package bruteforce_test

import (
	"slices"
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestReverseSlice(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "even length",
			nums: []int{1, 2, 3, 4},
			want: []int{4, 3, 2, 1},
		},
		{
			name: "odd length",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "single value",
			nums: []int{7},
			want: []int{7},
		},
		{
			name: "empty",
			nums: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bruteforce.ReverseSlice(tt.nums)

			if !slices.Equal(tt.nums, tt.want) {
				t.Fatalf("ReverseSlice() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
