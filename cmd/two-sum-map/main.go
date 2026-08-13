package main

import (
	"fmt"

	hashmap "github.com/Nextjingjing/go-algorithm-lab/algorithms/hash-map"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(hashmap.TwoSum(nums, target))
}
