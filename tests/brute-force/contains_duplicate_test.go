package bruteforce_test

import (
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "has duplicate",
			nums: []int{1, 2, 3, 1},
			want: true,
		},
		{
			name: "no duplicate",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "duplicate next to each other",
			nums: []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "single value",
			nums: []int{1},
			want: false,
		},
		{
			name: "empty",
			nums: []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteforce.ContainsDuplicate(tt.nums)

			if got != tt.want {
				t.Fatalf("ContainsDuplicate(%v) = %t, want %t", tt.nums, got, tt.want)
			}
		})
	}
}
