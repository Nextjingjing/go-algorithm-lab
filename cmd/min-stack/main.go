package main

import (
	"fmt"

	stack "github.com/Nextjingjing/go-algorithm-lab/algorithms/stack"
)

func main() {
	minStack := stack.Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	fmt.Println("minimum:", minStack.GetMin())
	minStack.Pop()
	fmt.Println("top:", minStack.Top())
	fmt.Println("minimum:", minStack.GetMin())
}
