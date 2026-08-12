package main

import (
	"fmt"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func main() {
	nums := []int{1, 2, 2, 3, 2}
	target := 2

	result := bruteforce.CountOccurrences(nums, target)

	fmt.Println(result)
}
