package main

import (
	"fmt"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func main() {
	nums := []int{10, 4, 8, 3}

	fmt.Printf("LeftRightDifference(%v) = %v\n", nums, prefixsuffix.LeftRightDifference(nums))
}
