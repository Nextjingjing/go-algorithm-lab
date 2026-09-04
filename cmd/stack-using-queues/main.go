package main

import (
	"fmt"

	stack "github.com/Nextjingjing/go-algorithm-lab/algorithms/stack"
)

func main() {
	s := stack.NewStackUsingQueues()
	s.Push(10)
	s.Push(20)

	fmt.Println("top:", s.Top())
	fmt.Println("pop:", s.Pop())
	fmt.Println("empty:", s.Empty())
}
