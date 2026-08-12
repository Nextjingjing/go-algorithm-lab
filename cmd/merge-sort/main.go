package main

import (
	"fmt"

	divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"
)

func main() {
	data := []int{8, 3, 5, 1, 9, 2}

	divideconquer.MergeSort(data)

	fmt.Println(data)
}
