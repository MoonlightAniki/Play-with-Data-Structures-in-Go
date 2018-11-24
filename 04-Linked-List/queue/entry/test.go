package main

import (
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/queue"
	"time"
	"math/rand"
	"fmt"
)

func testQueue(queue queue.Queue, opCount int) float64 {
	startTime := time.Now()
	for i := 0; i < opCount; i++ {
		queue.Offer(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		queue.Poll()
	}
	return time.Now().Sub(startTime).Seconds()
}

func main() {
	opCount := 100000

	linkedListQueue := queue.CreateLinkedListQueue()
	fmt.Println("LinkedListQueue, time:", testQueue(linkedListQueue, opCount), "s")

	loopQueue := queue.CreateLoopQueue()
	fmt.Println("LoopQueue, time:", testQueue(loopQueue, opCount), "s")
}
