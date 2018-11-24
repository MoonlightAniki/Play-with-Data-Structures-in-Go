package queue

import (
	"dev/Play-with-Data-Structures-in-Go/03-Stacks-and-Queues/array"
	"bytes"
	"fmt"
)

type ArrayQueue struct {
	arr *array.Array
}

func CreateArrayQueue() *ArrayQueue {
	queue := &ArrayQueue{}
	queue.arr = array.CreateArray()
	return queue
}

func (queue *ArrayQueue) GetSize() int {
	return queue.arr.GetSize()
}

func (queue *ArrayQueue) IsEmpty() bool {
	return queue.arr.IsEmpty()
}

func (queue *ArrayQueue) Offer(e interface{}) {
	queue.arr.AddLast(e)
}

func (queue *ArrayQueue) Poll() interface{} {
	return queue.arr.RemoveFirst()
}

func (queue *ArrayQueue) Front() interface{} {
	return queue.arr.GetFirst()
}

func (queue *ArrayQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("ArrayQueue, front [")
	for i := 0; i < queue.arr.GetSize(); i++ {
		buffer.WriteString(fmt.Sprintf("%v", queue.arr.Get(i)))
		if i != queue.arr.GetSize()-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
