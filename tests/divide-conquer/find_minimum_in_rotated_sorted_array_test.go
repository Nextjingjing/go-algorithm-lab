package divideconquer_test

import (
	"reflect"
	"testing"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "odd length rotated", nums: []int{3, 4, 5, 1, 2}, want: 1},
		{name: "minimum is zero", nums: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{name: "ascending order", nums: []int{11, 13, 15, 17}, want: 11},
		{name: "one value", nums: []int{8}, want: 8},
		{name: "two values ascending", nums: []int{1, 2}, want: 1},
		{name: "two values rotated", nums: []int{2, 1}, want: 1},
		{name: "minimum at last index", nums: []int{2, 3, 4, 5, 1}, want: 1},
		{name: "minimum at second index", nums: []int{5, 1, 2, 3, 4}, want: 1},
		{name: "six values minimum near end", nums: []int{3, 4, 5, 6, 1, 2}, want: 1},
		{name: "even length rotated", nums: []int{5, 7, 1, 3}, want: 1},
		{name: "negative minimum", nums: []int{0, 3, -7, -4}, want: -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := append([]int(nil), tt.nums...)
			got := divideconquer.FindMin(tt.nums)
			if got != tt.want {
				t.Errorf("FindMin(%v) = %d, want %d", before, got, tt.want)
			}
			if !reflect.DeepEqual(tt.nums, before) {
				t.Errorf("FindMin modified nums: got %v, want %v", tt.nums, before)
			}
		})
	}
}
