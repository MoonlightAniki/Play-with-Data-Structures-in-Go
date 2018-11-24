package stack

import (
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/list"
	"bytes"
	"fmt"
)

type LinkedListStack struct {
	list *list.LinkedList
}

func CreateLinkedListStack() *LinkedListStack {
	return &LinkedListStack{list.CreateLinkedList()}
}

func (stack *LinkedListStack) GetSize() int {
	return stack.list.GetSize()
}

func (stack *LinkedListStack) IsEmpty() bool {
	return stack.list.IsEmpty()
}

func (stack *LinkedListStack) Push(e interface{}) {
	stack.list.AddFirst(e)
}

func (stack *LinkedListStack) Pop() interface{} {
	return stack.list.RemoveFirst()
}

func (stack *LinkedListStack) Peek() interface{} {
	return stack.list.GetFirst()
}

func (stack *LinkedListStack) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LinkedListStack, size = %d, top [", stack.GetSize()))
	for i := 0; i < stack.list.GetSize(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", stack.list.Get(i)))
		if i != stack.list.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
