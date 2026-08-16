package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "buy low then sell high", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "prices only decrease", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "increasing prices", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "one price", prices: []int{5}, want: 0},
		{name: "empty prices", prices: []int{}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slidingwindow.MaxProfit(tt.prices)
			if got != tt.want {
				t.Fatalf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
