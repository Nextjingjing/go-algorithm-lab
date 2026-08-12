package main

import (
	"fmt"

	divideconquer "my-algorithm/algorithms/divide-conquer"
)

func main() {
	data := []int{8, 3, 5, 1, 9, 2}

	divideconquer.MergeSort(&data)

	fmt.Println(data)
}
