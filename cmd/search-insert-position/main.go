package main

import (
	"fmt"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	position := divideconquer.SearchInsert(nums, target)
	fmt.Printf("insert %d at index %d in %v\n", target, position, nums)
}
