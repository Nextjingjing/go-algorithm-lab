package bruteforce_test

import (
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestIsSorted(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "sorted",
			nums: []int{1, 2, 3, 4},
			want: true,
		},
		{
			name: "not sorted",
			nums: []int{1, 3, 2, 4},
			want: false,
		},
		{
			name: "allows equal neighbors",
			nums: []int{1, 2, 2, 3},
			want: true,
		},
		{
			name: "single value",
			nums: []int{1},
			want: true,
		},
		{
			name: "empty",
			nums: []int{},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteforce.IsSorted(tt.nums)

			if got != tt.want {
				t.Fatalf("IsSorted(%v) = %t, want %t", tt.nums, got, tt.want)
			}
		})
	}
}
