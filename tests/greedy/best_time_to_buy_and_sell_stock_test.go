package greedy_test

import (
	"testing"

	greedy "github.com/Nextjingjing/go-algorithm-lab/algorithms/greedy"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "buy low then sell high", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "prices only decrease", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "best transaction uses first and last day", prices: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "one price", prices: []int{5}, want: 0},
		{name: "empty prices", prices: []int{}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greedy.MaxProfit(tt.prices)
			if got != tt.want {
				t.Fatalf("MaxProfit(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
