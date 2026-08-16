package main

import (
	"fmt"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}

	fmt.Printf("MaxConsecutiveOnes(%v) = %d\n", nums, slidingwindow.MaxConsecutiveOnes(nums))
}
