package main

import (
	"fmt"

	twopointers "github.com/Nextjingjing/go-algorithm-lab/algorithms/two-pointers"
)

func main() {
	s := "A man, a plan, a canal: Panama"

	fmt.Printf("IsPalindrome(%q) = %t\n", s, twopointers.IsPalindrome(s))
}
