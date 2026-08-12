package main

import (
	"fmt"

	bruteforce "github.com/Nextjingjing/go-algorithm-lab/algorithms/brute-force"
)

func main() {
	nums := []int{4, 9, 1, 12, 7}

	result := bruteforce.FindMax(&nums)

	fmt.Println(result)
}
