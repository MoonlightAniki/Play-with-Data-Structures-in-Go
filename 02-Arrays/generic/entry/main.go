package main

import (
	"dev/Play-with-Data-Structures-in-Go/02-Arrays/generic"
	"fmt"
)

func main() {
	arr := generic.GetArray()

	fmt.Println(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())

	arr.AddLast(generic.GetStudent("Alice", 99))
	arr.AddLast(generic.GetStudent("Bob", 75))
	arr.AddFirst(generic.GetStudent("Robin", 100))
	fmt.Println(arr)
	fmt.Println(arr.GetCapacity(), arr.GetSize(), arr.IsEmpty())

	fmt.Println(arr.RemoveLast())



}
