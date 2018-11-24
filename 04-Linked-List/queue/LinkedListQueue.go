package queue

import (
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/list"
	"bytes"
	"fmt"
)

type LinkedListQueue struct {
	list *list.LinkedList
}

func CreateLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{list.CreateLinkedList()}
}

func (queue *LinkedListQueue) GetSize() int {
	return queue.list.GetSize()
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.list.IsEmpty()
}

func (queue *LinkedListQueue) Offer(e interface{}) {
	queue.list.AddLast(e)
}

func (queue *LinkedListQueue) Poll() interface{} {
	return queue.list.RemoveFirst()
}

func (queue *LinkedListQueue) Front() interface{} {
	return queue.list.GetFirst()
}

func (queue *LinkedListQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LinkedListQueue, size = %d, front[", queue.list.GetSize()))
	for i := 0; i < queue.list.GetSize(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", queue.list.Get(i)))
		if i != queue.list.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]tail")
	return buffer.String()
}
