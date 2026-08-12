package bruteforce_test

import (
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestSumArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "positive numbers",
			nums: []int{1, 2, 3},
			want: 6,
		},
		{
			name: "empty slice",
			nums: []int{},
			want: 0,
		},
		{
			name: "negative numbers",
			nums: []int{-2, 5, -1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteforce.SumArray(tt.nums)

			if got != tt.want {
				t.Fatalf("SumArray(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
