package main

import (
	"fmt"

	greedy "github.com/Nextjingjing/go-algorithm-lab/algorithms/greedy"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Printf("MaxProfit(%v) = %d\n", prices, greedy.MaxProfit(prices))
}
