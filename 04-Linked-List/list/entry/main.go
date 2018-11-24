package main

import (
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/list"
	"fmt"
)

func main() {
	l := list.CreateLinkedList()
	for i := 0; i < 5; i++ {
		l.AddLast(i)
		fmt.Println(l)
	}
	for i := 5; i < 10; i++ {
		l.AddFirst(i)
		fmt.Println(l)
	}

	l.Set(4, 100)
	fmt.Println(l)

	fmt.Println(l.RemoveFirst())
	fmt.Println(l.RemoveLast())
	fmt.Println(l)
}
