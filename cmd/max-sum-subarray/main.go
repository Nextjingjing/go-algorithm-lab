package main

import (
	"fmt"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3

	fmt.Printf("MaxSumSubarray(%v, %d) = %d\n", nums, k, slidingwindow.MaxSumSubarray(nums, k))
}
