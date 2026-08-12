package bruteforce_test

import (
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestCountOccurrences(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "found multiple times",
			nums:   []int{1, 2, 2, 3, 2},
			target: 2,
			want:   3,
		},
		{
			name:   "found once",
			nums:   []int{4, 8, 2, 9},
			target: 8,
			want:   1,
		},
		{
			name:   "not found",
			nums:   []int{4, 8, 2, 9},
			target: 7,
			want:   0,
		},
		{
			name:   "empty",
			nums:   []int{},
			target: 1,
			want:   0,
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -1, 3},
			target: -1,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteforce.CountOccurrences(tt.nums, tt.target)

			if got != tt.want {
				t.Fatalf("CountOccurrences(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
