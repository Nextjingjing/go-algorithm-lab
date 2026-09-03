package main

import (
	"fmt"

	queue "github.com/Nextjingjing/go-algorithm-lab/algorithms/queue"
)

func main() {
	q := queue.Constructor()
	q.Push(1)
	q.Push(2)

	fmt.Println("front:", q.Peek())
	fmt.Println("pop:", q.Pop())
	fmt.Println("empty:", q.Empty())
}
