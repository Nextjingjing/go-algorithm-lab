package main

import (
	"fmt"

	stack "github.com/Nextjingjing/go-algorithm-lab/algorithms/stack"
)

func main() {
	s := "{[]}"
	fmt.Printf("IsValid(%q) = %t\n", s, stack.IsValid(s))
}
