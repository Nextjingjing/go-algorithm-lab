package main

import (
	"fmt"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Printf("SumArray(%v) = %d\n", nums, bruteforce.SumArray(nums))
}
