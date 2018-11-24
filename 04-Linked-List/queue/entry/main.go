package main

import (
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/queue"
	"math/rand"
	"fmt"
)

func main() {
	llq := queue.CreateLinkedListQueue()
	for i := 0; i < 10; i++ {
		llq.Offer(rand.Int() % 100)
	}
	fmt.Println(llq)

	for !llq.IsEmpty() {
		fmt.Println(llq.Front())
		fmt.Println(llq.Poll())
		fmt.Println(llq)
	}

	fmt.Println("-------------------------------")
	loopQueue := queue.CreateLoopQueue()
	for i := 0; i < 10; i++ {
		loopQueue.Offer(rand.Int() % 100)
	}
	fmt.Println(loopQueue)

	for !loopQueue.IsEmpty() {
		fmt.Println(loopQueue.Front())
		fmt.Println(loopQueue.Poll())
		fmt.Println(loopQueue)
	}
}
