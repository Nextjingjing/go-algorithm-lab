package prefixsuffix_test

import (
	"testing"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "pivot in the middle",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "no pivot",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "pivot at the first index",
			nums: []int{2, 1, -1},
			want: 0,
		},
		{
			name: "pivot at the last index",
			nums: []int{-1, 1, 2},
			want: 2,
		},
		{
			name: "single value",
			nums: []int{5},
			want: 0,
		},
		{
			name: "empty slice",
			nums: []int{},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixsuffix.PivotIndex(tt.nums)

			if got != tt.want {
				t.Fatalf("PivotIndex(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
