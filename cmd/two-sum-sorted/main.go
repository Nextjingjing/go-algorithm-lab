package main

import (
	"fmt"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("TwoSumSorted(%v, %d) = %v\n", numbers, target, twopointers.TwoSumSorted(numbers, target))
}
