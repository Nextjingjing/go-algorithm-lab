package main

import (
	"fmt"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	k := twopointers.RemoveElement(nums, val)

	fmt.Printf("RemoveElement: k=%d, nums[:k]=%v\n", k, nums[:k])
}
