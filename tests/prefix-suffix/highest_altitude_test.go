package prefixsuffix_test

import (
	"testing"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func TestLargestAltitude(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{name: "highest point in the middle", gain: []int{-5, 1, 5, 0, -7}, want: 1},
		{name: "never rises above starting altitude", gain: []int{-4, -3, -2, -1, 4, 3, 2}, want: 0},
		{name: "one gain", gain: []int{0}, want: 0},
		{name: "empty gains", gain: []int{}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixsuffix.LargestAltitude(tt.gain)
			if got != tt.want {
				t.Fatalf("LargestAltitude(%v) = %d, want %d", tt.gain, got, tt.want)
			}
		})
	}
}
