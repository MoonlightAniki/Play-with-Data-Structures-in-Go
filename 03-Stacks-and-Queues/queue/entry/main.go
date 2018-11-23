package main

import (
	"dev/Play-with-Data-Structures-in-Go/03-Stacks-and-Queues/queue"
	"fmt"
)

func main() {
	queue := queue.CreateArrayQueue()
	fmt.Println(queue)

	queue.Offer(1)
	queue.Offer(2)
	queue.Offer(3)
	fmt.Println(queue)

	fmt.Println(queue.Front())
	fmt.Println(queue.Poll())
	fmt.Println(queue)
}
