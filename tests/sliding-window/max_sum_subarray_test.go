package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func TestMaxSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "largest sum in a window", nums: []int{2, 1, 5, 1, 3, 2}, k: 3, want: 9},
		{name: "window of one", nums: []int{1, 2, 3}, k: 1, want: 3},
		{name: "window covers all values", nums: []int{1, 2, 3}, k: 3, want: 6},
		{name: "negative values", nums: []int{-2, -1, -3}, k: 2, want: -3},
		{name: "empty input", nums: []int{}, k: 1, want: 0},
		{name: "invalid window size", nums: []int{1, 2, 3}, k: 0, want: 0},
		{name: "window larger than input", nums: []int{1, 2, 3}, k: 4, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slidingwindow.MaxSumSubarray(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("MaxSumSubarray(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
