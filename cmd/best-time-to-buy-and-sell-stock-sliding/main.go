package main

import (
	"fmt"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Printf("MaxProfit(%v) = %d\n", prices, slidingwindow.MaxProfit(prices))
}
