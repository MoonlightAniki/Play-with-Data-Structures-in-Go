package main

import (
	"fmt"
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/stack"
	"math/rand"
)

func main() {
	stack := stack.CreateLinkedListStack()
	fmt.Println(stack)

	for i := 0; i < 20; i++ {
		stack.Push(rand.Int() % 1000)
	}
	fmt.Println(stack)

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
		fmt.Println(stack)
	}
}
