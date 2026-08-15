package prefixsuffix_test

import (
	"slices"
	"testing"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "positive numbers",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "negative and zero values",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
		{
			name: "one zero",
			nums: []int{1, 2, 0, 4},
			want: []int{0, 0, 8, 0},
		},
		{
			name: "multiple zeros",
			nums: []int{0, 1, 0, 3},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "two values",
			nums: []int{3, 5},
			want: []int{5, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixsuffix.ProductExceptSelf(tt.nums)

			if !slices.Equal(got, tt.want) {
				t.Fatalf("ProductExceptSelf(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
