package divideconquer_test

import (
	"testing"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name string
		data []int
		x    int
		want int
	}{
		{
			name: "found in middle",
			data: []int{1, 3, 5, 7, 9},
			x:    5,
			want: 2,
		},
		{
			name: "found at left",
			data: []int{1, 3, 5, 7, 9},
			x:    1,
			want: 0,
		},
		{
			name: "not found",
			data: []int{1, 3, 5, 7, 9},
			x:    0,
			want: -1,
		},
		{
			name: "empty slice",
			data: []int{},
			x:    2,
			want: -1,
		},
		{
			name: "x is less than slice's minimum value",
			data: []int{1, 3, 5, 7, 9},
			x:    0,
			want: -1,
		},
		{
			name: "x is more than slice's maximum value",
			data: []int{1, 3, 5, 7, 9},
			x:    10,
			want: -1,
		},
		{
			name: "one element",
			data: []int{1},
			x:    1,
			want: 0,
		},
		{
			name: "found at last",
			data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
			x:    2,
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divideconquer.BinarySearch(tt.data, tt.x)

			if got != tt.want {
				t.Fatalf("BinarySearch(%v, %d) = %d, want %d", tt.data, tt.x, got, tt.want)
			}
		})
	}
}
