package main

import (
	"fmt"

	prefixsuffix "github.com/Nextjingjing/go-algorithm-lab/algorithms/prefix-suffix"
)

func main() {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{name: "pivot in the middle", nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{name: "pivot at the first index", nums: []int{2, 1, -1}, want: 0},
	}

	for _, tc := range cases {
		got := prefixsuffix.PivotIndex(tc.nums)
		fmt.Printf("%s: nums=%v, got=%d, want=%d\n", tc.name, tc.nums, got, tc.want)
	}
}
