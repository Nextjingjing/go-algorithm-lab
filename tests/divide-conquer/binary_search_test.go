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
