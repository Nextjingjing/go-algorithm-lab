package main

import (
	"fmt"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	minimum := divideconquer.FindMin(nums)
	fmt.Printf("minimum in %v = %d\n", nums, minimum)
}
