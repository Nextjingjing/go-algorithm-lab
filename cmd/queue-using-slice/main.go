package main

import (
	"fmt"

	queue "github.com/Nextjingjing/go-algorithm-lab/algorithms/queue"
)

func main() {
	q := queue.NewQueueUsingSlice()
	q.Enqueue(10)
	q.Enqueue(20)

	fmt.Println("front:", q.Peek())
	fmt.Println("dequeue:", q.Dequeue())
	fmt.Println("empty:", q.Empty())
}
