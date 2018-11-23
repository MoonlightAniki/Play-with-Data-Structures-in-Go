package main

import (
	"dev/Play-with-Data-Structures-in-Go/03-Stacks-and-Queues/stack"
	"fmt"
)

func main() {
	stack := stack.CreateArrayStack()
	fmt.Println(stack)
	fmt.Println(stack.GetSize(), stack.IsEmtpy())

	for i := 0; i < 20; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(stack.Peek())
		fmt.Println(stack.Pop())
		fmt.Println(stack)
	}
}
