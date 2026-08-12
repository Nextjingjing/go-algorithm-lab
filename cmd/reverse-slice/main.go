package main

import (
	"fmt"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func main() {
	nums := []int{1, 2, 3, 4}

	bruteforce.ReverseSlice(nums)

	fmt.Println(nums)
}
