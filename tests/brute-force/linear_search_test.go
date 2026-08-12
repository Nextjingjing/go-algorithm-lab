package bruteforce_test

import (
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "found in middle",
			nums:   []int{4, 8, 2, 9},
			target: 2,
			want:   2,
		},
		{
			name:   "found first",
			nums:   []int{4, 8, 2, 9},
			target: 4,
			want:   0,
		},
		{
			name:   "found last",
			nums:   []int{4, 8, 2, 9},
			target: 9,
			want:   3,
		},
		{
			name:   "not found",
			nums:   []int{4, 8, 2, 9},
			target: 7,
			want:   -1,
		},
		{
			name:   "empty",
			nums:   []int{},
			target: 1,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteforce.LinearSearch(tt.nums, tt.target)

			if got != tt.want {
				t.Fatalf("LinearSearch(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
