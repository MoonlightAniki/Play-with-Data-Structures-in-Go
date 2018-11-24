package main

import (
	"dev/Play-with-Data-Structures-in-Go/05-Recursion/list"
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

	fmt.Println(l.Get(5))

	for i := 0; i < 20; i++ {
		fmt.Println(l.Contains(i))
	}

	l.Set(5, 100)
	fmt.Println(l)

	fmt.Println(l.Remove(5))
	fmt.Println(l, l.GetSize())

	//for !l.IsEmpty() {
	//	fmt.Println(l.RemoveLast())
	//	fmt.Println(l)
	//}

	//fmt.Println(l.RemoveElement(1))
	//fmt.Println(l, l.GetSize())
	//
	//for i := 2; i < 10; i++ {
	//	l.RemoveElement(i)
	//	fmt.Println(l, l.GetSize())
	//}
	l.AddFirst(9)
	l.AddLast(9)
	l.Add(6, 9)
	fmt.Println(l, l.GetSize())

	fmt.Println(l.RemoveAll(9))
	fmt.Println(l, l.GetSize())

	fmt.Println(l.RemoveAll(100))
	fmt.Println(l, l.GetSize())

	fmt.Println(l.RemoveAll(1))
	fmt.Println(l, l.GetSize())

}
