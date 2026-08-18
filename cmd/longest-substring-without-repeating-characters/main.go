package main

import (
	"fmt"

	slidingwindow "github.com/Nextjingjing/go-algorithm-lab/algorithms/sliding-window"
)

func main() {
	s := "abcbacbb"

	fmt.Printf("LengthOfLongestSubstring(%q) = %d\n", s, slidingwindow.LengthOfLongestSubstring(s))
}
