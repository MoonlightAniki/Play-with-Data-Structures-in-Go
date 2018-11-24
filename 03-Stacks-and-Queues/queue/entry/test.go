package main

import (
	"dev/Play-with-Data-Structures-in-Go/03-Stacks-and-Queues/queue"
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

	arrayQueue := queue.CreateArrayQueue()
	time := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time, "s")

	loopQueue := queue.CreateLoopQueue()
	time = testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time, "s")
}

//ArrayQueue, time: 6.951404 s
//LoopQueue, time: 0.015984 s
