package main

import (
	"dev/Play-with-Data-Structures-in-Go/02-Arrays/array"
	"fmt"
)

func main() {
	a := array.GetArray(5)

	fmt.Println(a)
	fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())

	for i := 0; i < 5; i++ {
		a.AddLast(i)
	}
	fmt.Println(a)
	fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())

	for i := 4; i >= 0; i-- {
		fmt.Println(a.Get(i))
	}

	for i := 0; i < 5; i++ {
		a.Set(i, a.Get(i)*a.Get(i))
	}
	fmt.Println(a)

	fmt.Println(a.Contains(9))
	fmt.Println(a.Contains(7))

	fmt.Println(a.Find(15))
	fmt.Println(a.Find(16))

	fmt.Println(a.FindAll(5))
	fmt.Println(a.FindAll(4))

	fmt.Println(a.Remove(4))
	fmt.Println(a)

	fmt.Println(a.RemoveElement(2))
	fmt.Println(a.RemoveElement(4))
	fmt.Println(a)

	a.AddFirst(1)
	a.AddFirst(1)
	fmt.Println(a.RemoveAll(1))
	fmt.Println(a)
}
