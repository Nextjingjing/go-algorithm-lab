package main

import divideconquer "github.com/Nextjingjing/go-algorithm-lab/algorithms/divide-conquer"

func main() {
	d := []int{2, 4, 6, 7, 8, 9, 11}
	i := divideconquer.BinarySearch(d, 0)
	print(i)
}
