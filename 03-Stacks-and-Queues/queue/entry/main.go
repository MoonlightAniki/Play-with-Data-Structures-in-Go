package main

import (
	"dev/Play-with-Data-Structures-in-Go/03-Stacks-and-Queues/queue"
	"fmt"
)

func main() {
	aq := queue.CreateArrayQueue()
	fmt.Println(aq)

	aq.Offer(1)
	aq.Offer(2)
	aq.Offer(3)
	fmt.Println(aq)

	fmt.Println(aq.Front())
	fmt.Println(aq.Poll())
	fmt.Println(aq)

	fmt.Println("-----------------------------")

	lq := queue.CreateLoopQueue()
	fmt.Println(lq)

	for i := 0; i < 20; i++ {
		lq.Offer(i)
		fmt.Println(lq)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(lq.Front())
		fmt.Println(lq.Poll())
		fmt.Println(lq)
	}
}
