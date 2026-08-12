package bruteforce_test

import (
	"testing"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func TestFindMax(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "positive numbers",
			data: []int{4, 9, 1, 12, 7},
			want: 12,
		},
		{
			name: "single value",
			data: []int{5},
			want: 5,
		},
		{
			name: "negative numbers",
			data: []int{-8, -3, -20, -1},
			want: -1,
		},
		{
			name: "duplicates",
			data: []int{3, 9, 9, 2},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bruteforce.FindMax(&tt.data)

			if got != tt.want {
				t.Fatalf("FindMax(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
