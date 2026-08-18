package main

import (
	"fmt"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func main() {
	gain := []int{-5, 1, 5, 0, -7}

	fmt.Printf("LargestAltitude(%v) = %d\n", gain, prefixsuffix.LargestAltitude(gain))
}
