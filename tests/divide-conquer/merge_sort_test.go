package divideconquer_test

import (
	"slices"
	"testing"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{
			name: "unsorted",
			data: []int{8, 3, 5, 1, 9, 2},
			want: []int{1, 2, 3, 5, 8, 9},
		},
		{
			name: "already sorted",
			data: []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "reverse sorted",
			data: []int{4, 3, 2, 1},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "duplicates",
			data: []int{3, 1, 2, 3, 1},
			want: []int{1, 1, 2, 3, 3},
		},
		{
			name: "empty",
			data: []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			divideconquer.MergeSort(tt.data)

			if !slices.Equal(tt.data, tt.want) {
				t.Fatalf("MergeSort() = %v, want %v", tt.data, tt.want)
			}
		})
	}
}
