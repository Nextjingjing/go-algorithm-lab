package main

import (
	"fmt"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func main() {
	nums := []int{4, 8, 2, 9}
	target := 2

	result := bruteforce.LinearSearch(nums, target)

	fmt.Println(result)
}
