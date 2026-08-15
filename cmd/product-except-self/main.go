package main

import (
	"fmt"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func main() {
	nums := []int{2, 3, 4, 5}

	fmt.Printf("ProductExceptSelf(%v) = %v\n", nums, prefixsuffix.ProductExceptSelf(nums))
}
