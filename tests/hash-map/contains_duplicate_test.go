package hashmap_test

import (
	"testing"

	hashmap "github.com/Nextjingjing/go-algorithm-lab/algorithms/hash-map"
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
			name: "has no duplicate",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "empty slice",
			nums: []int{},
			want: false,
		},
		{
			name: "single value",
			nums: []int{7},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashmap.ContainsDuplicate(tt.nums)

			if got != tt.want {
				t.Fatalf("ContainsDuplicate(%v) = %t, want %t", tt.nums, got, tt.want)
			}
		})
	}
}
