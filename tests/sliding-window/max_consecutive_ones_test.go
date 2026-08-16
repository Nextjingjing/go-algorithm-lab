package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func TestMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "longest run in the middle", nums: []int{1, 1, 0, 1, 1, 1}, want: 3},
		{name: "runs separated by zeroes", nums: []int{1, 0, 1, 1, 0, 1}, want: 2},
		{name: "all zeroes", nums: []int{0, 0, 0}, want: 0},
		{name: "all ones", nums: []int{1, 1, 1, 1}, want: 4},
		{name: "run after a zero", nums: []int{1, 0, 1, 1}, want: 2},
		{name: "one after multiple zeroes", nums: []int{0, 0, 1}, want: 1},
		{name: "empty input", nums: []int{}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slidingwindow.MaxConsecutiveOnes(tt.nums)
			if got != tt.want {
				t.Fatalf("MaxConsecutiveOnes(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
