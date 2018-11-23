package stack

import (
	"dev/Play-with-Data-Structures-in-Go/03-Stacks-and-Queues/array"
	"bytes"
	"fmt"
)

type ArrayStack struct {
	arr *array.Array
}

func CreateArrayStack() *ArrayStack {
	stack := &ArrayStack{}
	stack.arr = array.CreateArray()
	return stack
}

func (stack *ArrayStack) GetSize() int {
	return stack.arr.GetSize()
}

func (stack *ArrayStack) IsEmtpy() bool {
	return stack.arr.IsEmpty()
}

func (stack *ArrayStack) Push(e interface{}) {
	stack.arr.AddLast(e)
}

func (stack *ArrayStack) Pop() interface{} {
	return stack.arr.RemoveLast()
}

func (stack *ArrayStack) Peek() interface{} {
	return stack.arr.GetLast()
}

func (stack *ArrayStack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Stack [")
	for i := 0; i < stack.arr.GetSize(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", stack.arr.Get(i)))
		if i != stack.arr.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] top")
	return buffer.String()
}
