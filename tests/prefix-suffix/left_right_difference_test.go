package prefixsuffix_test

import (
	"slices"
	"testing"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func TestLeftRightDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "positive values", nums: []int{10, 4, 8, 3}, want: []int{15, 1, 11, 22}},
		{name: "mixed values", nums: []int{1, 2, 3, 4, 5}, want: []int{14, 11, 6, 1, 10}},
		{name: "two values", nums: []int{5, 1}, want: []int{1, 5}},
		{name: "one value", nums: []int{7}, want: []int{0}},
		{name: "empty input", nums: []int{}, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixsuffix.LeftRightDifference(tt.nums)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("LeftRightDifference(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}
