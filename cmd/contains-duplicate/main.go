package main

import (
	"fmt"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func main() {
	nums := []int{1, 2, 3, 1}

	result := bruteforce.ContainsDuplicate(nums)

	fmt.Println(result)
}
