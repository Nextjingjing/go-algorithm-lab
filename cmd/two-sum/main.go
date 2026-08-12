package main

import (
	"fmt"

	bruteforce "my-algorithm/algorithms/brute-force"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := bruteforce.TwoSum(nums, target)

	fmt.Println(result)
}
