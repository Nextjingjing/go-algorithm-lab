package main

import (
	"fmt"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func main() {
	nums := []int{0, 1, 0, 3, 12}
	twopointers.MoveZeroes(nums)

	fmt.Printf("MoveZeroes = %v\n", nums)
}
