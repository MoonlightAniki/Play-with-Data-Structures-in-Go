package queue

import (
	"bytes"
	"fmt"
)

type node struct {
	val  interface{}
	next *node
}

type linkedList struct {
	front *node
	tail  *node
	size  int
}

func (list *linkedList) getSize() int {
	return list.size
}

func (list *linkedList) isEmpty() bool {
	return list.size == 0
}

func createLinkedList() *linkedList {
	list := &linkedList{}
	list.front = nil
	list.tail = nil
	list.size = 0
	return list
}

func (list *linkedList) addLast(val interface{}) {
	if list.front == nil {
		list.front = &node{val, nil}
		list.tail = list.front
		list.size++
	} else {
		list.tail.next = &node{val, nil}
		list.tail = list.tail.next
		list.size++
	}
}

func (list *linkedList) removeFirst() interface{} {
	if list.size == 0 {
		panic("list is empty")
	}
	front := list.front
	list.front = front.next
	front.next = nil
	list.size--
	return front.val
}

func (list *linkedList) getFirst() interface{} {
	if list.size == 0 {
		panic("list is empty")
	}
	return list.front.val
}

type LoopQueue struct {
	list *linkedList
}

func CreateLoopQueue() *LoopQueue {
	return &LoopQueue{createLinkedList()}
}

func (queue *LoopQueue) GetSize() int {
	return queue.list.getSize()
}

func (queue *LoopQueue) IsEmpty() bool {
	return queue.list.isEmpty()
}

func (queue *LoopQueue) Offer(e interface{}) {
	queue.list.addLast(e)
}

func (queue *LoopQueue) Poll() interface{} {
	return queue.list.removeFirst()
}

func (queue *LoopQueue) Front() interface{} {
	return queue.list.getFirst()
}

func (queue *LoopQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("LoopQueue, size = %d, front[", queue.list.getSize()))
	for cur := queue.list.front; cur != nil; cur = cur.next {
		buffer.WriteString(fmt.Sprintf("%v", cur.val))
		if cur != queue.list.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]tail")
	return buffer.String()
}
